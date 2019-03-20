package mydsl

import (
	"context"
	_ "fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	//	"reflect"
	"os"
	"regexp"
	"time"
)

var mongoDbnamePattern = regexp.MustCompile(`^mongodb://(.+?):`)

func init() {
	mongodbUri := os.Getenv("MONGODB_URI")
	dbname := mongoDbnamePattern.FindStringSubmatch(mongodbUri)[1]
	uriOption := options.Client().ApplyURI(mongodbUri)
	client, _ := mongo.NewClient(uriOption)
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Hour)
	client.Connect(ctx)
	MongoCollection := func(collectionName string) *mongo.Collection {
		return client.Database(dbname).Collection(collectionName)
	}

	DslFunctions["mongoGet"] = func(container *map[string]interface{}, args ...Argument) (interface{}, error) {
		collectionName := args[0].RawArg.(string)
		collection := MongoCollection(collectionName)
		cur, err := collection.Find(ctx, bson.D{})
		if err != nil {
			log.Fatal(err)
		}
		records := []map[string]interface{}{}
		defer cur.Close(ctx)
		for cur.Next(ctx) {
			var result map[string]interface{}
			err := cur.Decode(&result)
			if err != nil {
				log.Fatal(err)
			}
			records = append(records, result)
		}
		if err := cur.Err(); err != nil {
			log.Fatal(err)
		}
		return records, nil
	}

	DslFunctions["mongoInsert"] = func(container *map[string]interface{}, args ...Argument) (interface{}, error) {
		collectionName := args[0].RawArg.(string)
		obj, err := args[1].Evaluate(container)
		if err != nil {
			return nil, err
		}
		collection := MongoCollection(collectionName)
		res, err := collection.InsertOne(ctx, obj)
		if err != nil {
			return nil, err
		}
		return res, nil
	}

	DslFunctions["mongoReplace"] = func(container *map[string]interface{}, args ...Argument) (interface{}, error) {
		collectionName := args[0].RawArg.(string)
		obj, err := args[1].Evaluate(container)
		if err != nil {
			return nil, err
		}
		collection := MongoCollection(collectionName)
		res := collection.FindOneAndReplace(ctx, map[string]interface{}{"_id": (obj.(map[string]interface{}))["_id"]}, obj)
		return res, nil
	}
}
