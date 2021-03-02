package main

import (
	"context"
	"fmt"
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"os"
	"time"
)

func setUpLogs() {
	log.SetFormatter(&log.JSONFormatter{})
	file, err := os.OpenFile("../config/logs.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		fmt.Println("No se ha podido crear el fichero de logs")
	}

	defer file.Close()
	log.SetOutput(file)
}

func goDotEnvVariable(key string) string {

	// load .env file
	err := godotenv.Load("config/.env")

	if err != nil {
		log.Error("Error loading .env file")
	}

	return os.Getenv(key)
}


func connectToMongo() *mongo.Collection {
	client, err := mongo.NewClient(
		options.Client().ApplyURI(
			goDotEnvVariable("MONGO_URI"),
		),
	)
	if err != nil { log.Error(err) }
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	err = client.Connect(ctx)
	if err != nil { log.Error(err) }
	collection := client.Database("Flow").Collection("activities")
	return collection
}

func getAllActivities(collection *mongo.Collection) {
	cur, err := collection.Find(context.Background(), bson.D{})
	if err != nil { log.Fatal(err) }
	defer cur.Close(context.Background())
	for cur.Next(context.Background()) {
		// To decode into a struct, use cursor.Decode()
		var act Activity
		err := cur.Decode(&act)
		if err != nil { log.Fatal(err) }
		// do something with result...
		fmt.Println(act.String())
		// To get the raw bson bytes use cursor.Current
		raw := cur.Current
		// do something with raw...
		fmt.Println(raw)
	}
	if err := cur.Err(); err != nil {
		log.Error(err)
	}
}