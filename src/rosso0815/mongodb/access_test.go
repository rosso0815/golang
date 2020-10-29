package main

import (
	//"context"
	"log"
	"testing"
	//"github.com/mongodb/mongo-go-driver/bson"
	//"github.com/mongodb/mongo-go-driver/mongo"
	//"github.com/mongodb/mongo-go-driver/mongo/options"
	//"go.mongodb.org/mongo-driver/mongo"

	"github.com/globalsign/mgo"
)

// You will be using this Trainer type later in the program
type Trainer struct {
	Name string
	Age  int
	City string
}

func TestConnection(t *testing.T) {

	log.Println("connectTest")
	_, err := mgo.Dial("test:test@localhost:27017/test")
	if err != nil {
		panic(err)
	}

	/*
		client, err := mongo.Connect(context.TODO(), "mongodb://localhost:27017")

		if err != nil {
			log.Fatal(err)
		}

		// Check the connection
		err = client.Ping(context.TODO(), nil)

		if err != nil {
			log.Fatal(err)
		}

		log.Println("Connected to MongoDB!")

		collection := client.Database("test").Collection("trainers")
		log.Println("collection=", collection)
		err = client.Disconnect(context.TODO())

		if err != nil {
			log.Fatal(err)
		}
		log.Println("Connection to MongoDB closed.")
	*/
}
