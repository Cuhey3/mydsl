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

type mongoUtil struct {
	Collection func(collectionName string) *mongo.Collection
	Context    context.Context
}

var sharedInstance *mongoUtil = newMongoUtil()

func newMongoUtil() *mongoUtil {
	mongodbUri := os.Getenv("MONGODB_URI")
	if mongodbUri == "" {
		return nil
	}
	dbname := regexp.MustCompile(`^mongodb://(.+?):`).FindStringSubmatch(mongodbUri)[1]
	uriOption := options.Client().ApplyURI(mongodbUri)
	client, _ := mongo.NewClient(uriOption)
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Hour)
	client.Connect(ctx)
	return &mongoUtil{func(collectionName string) *mongo.Collection {
		return client.Database(dbname).Collection(collectionName)
	}, ctx}
}

func MongoUtil() *mongoUtil {
	return sharedInstance
}

func init() {
	mongodbUri := os.Getenv("MONGODB_URI")
	if mongodbUri == "" {
		return
	}
	DslFunctions["mongoGet"] = func(container *map[string]interface{}, args ...Argument) (interface{}, error) {
		collectionName := args[0].RawArg.(string)
		collection := MongoUtil().Collection(collectionName)
		cur, err := collection.Find(MongoUtil().Context, bson.D{})
		if err != nil {
			log.Fatal(err)
		}
		records := []map[string]interface{}{}
		defer cur.Close(MongoUtil().Context)
		for cur.Next(MongoUtil().Context) {
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
		collection := MongoUtil().Collection(collectionName)
		res, err := collection.InsertOne(MongoUtil().Context, obj)
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
		collection := MongoUtil().Collection(collectionName)
		res := collection.FindOneAndReplace(MongoUtil().Context, map[string]interface{}{"_id": (obj.(map[string]interface{}))["_id"]}, obj)
		return res, nil
	}
}
