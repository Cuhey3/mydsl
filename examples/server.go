package main

import (
	"fmt"
	"log"
	"mydsl/go"
	"os"
)

func main() {
	yamlInput, err := mydsl.LoadYaml("examples/server.yml")
	if err != nil {
		log.Fatal(err)
	}
	port := os.Getenv("PORT")
	if port == "" {
		port = ":8080"
	} else {
		port = ":" + port
	}
	evaluated, err := mydsl.NewArgument(yamlInput).Evaluate(&(map[string]interface{}{"PORT": port}))
	fmt.Println(evaluated, err)
}
