package database

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"go.mongodb.org/mongo-driver/v2/mongo/readpref"
)

func MustGetMongoClient() (*mongo.Client){

	// Steps to connect
	// load env
	// make a client (mongo client) #check for errors
	// check the ping (ping makes shore that the connection is made properly)
	// at the end disconnect

	if err := godotenv.Load(".env"); err != nil {
		log.Fatal("there is an error loading env")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	clientOptions := options.Client().ApplyURI(os.Getenv("MONGO_CONNECTION_STRING"))
	client, err := mongo.Connect(clientOptions)

	if err != nil{
		log.Fatal("there is an error in connecting with database : ", err)
	}

	if err := client.Ping(ctx, readpref.Primary()); err != nil{
		log.Fatal("there is an error in connecting with database : ", err)
	}

	log.Println("successfully connected to the database")
	return client
}

// function to get a collection fromn mongoDB

func CreateMongCollection(client *mongo.Client, databaseName string, collectionName string) (*mongo.Collection){
	// creating a collection in database
	collection := client.Database(databaseName).Collection(collectionName) 
	return collection
}