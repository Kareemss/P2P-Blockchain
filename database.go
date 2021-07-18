package main

import (
	"context"
	"github.com/davecgh/go-spew/spew"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	_ "log"
	"os"
	"time"
)

func connectToDb() *mongo.Database {
	// Replace the uri string with your MongoDB deployment's connection string.
	username := os.Getenv("DB_USERNAME")
	password := os.Getenv("DB_PASSWORD")
	cluster := os.Getenv("DB_CLUSTER_ADDR")
	uri := "mongodb+srv://" + username + ":" + password + "@" + cluster + ".bzh1l.mongodb.net/myFirstDatabase?retryWrites=true&w=majority"
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}

	// Ping the primary
	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		panic(err)
	}
	spew.Dump("Successfully connected and pinged.")

	database := client.Database("Blockchain")

	return database
}

func addBlock(block Block, database *mongo.Database) *mongo.InsertOneResult {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	blocksCollection := database.Collection("blocks")

	insertionResult, err := blocksCollection.InsertOne(ctx, block)
	if err != nil {
		panic(err)
	}

	return insertionResult
}