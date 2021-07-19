package db

import (
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"context"
	"log"
)

var Client *mongo.Client

func InitDatabase() {
  var err error
  fmt.Println(MongoURI)
  Client, err = mongo.Connect(context.Background(), options.Client().ApplyURI(MongoURI))

  if err != nil {
    log.Fatal(err)
  }

  fmt.Println("Connected")
}

func Collection(collectionName string) (*mongo.Collection) {
  return Client.Database("imo").Collection(collectionName)
}