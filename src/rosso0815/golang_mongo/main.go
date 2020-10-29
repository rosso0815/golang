package main

import (
	"fmt"

	"github.com/mongodb/mongo-go-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo"
)

func main() {
	fmt.Println("Hello, world.")
	client, err := mongo.NewClient(options.Client().
	.ApplyURI("mongodb://localhost:27017"))
}
