package main

import (
	"context"
	"fmt"
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	. "go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	. "os"
	"time"
)

func setUpLogs() *File {
	log.SetFormatter(&log.JSONFormatter{})
	file, err := OpenFile("../config/logs.log", O_CREATE|O_WRONLY, 0666)
	if err != nil {
		fmt.Println("No se ha podido crear el fichero de logs")
	}

	log.SetOutput(file)
	return file
}

func closeLogFile(file *File) {
	file.Close()
}

func goDotEnvVariable(key string) string {

	// load .env file
	err := godotenv.Load("../config/.env")

	if err != nil {
		log.Error("Error loading .env file")
	}

	return Getenv(key)
}

func connectToMongo() (*Client, *Collection) {
	client, err := NewClient(
		options.Client().ApplyURI(
			goDotEnvVariable("MONGO_URI"),
		),
	)
	if err != nil { log.Error(err) }
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	err = client.Connect(ctx)
	if err != nil { log.Error(err) }
	log.Info("Conectado a mongo")
	collection := client.Database(goDotEnvVariable("DB")).Collection(goDotEnvVariable("COL"))
	log.Info("Elegida la base de datos y la colección")
	return client, collection
}

func getAllActivities(collection *Collection) {
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

func insertActivity(collection *Collection, act Activity) (bool)  {
	salida := true
	res, err := collection.InsertOne(context.Background(), act)
	if err != nil { log.Error(err) }
	id := res.InsertedID
	log.Info(fmt.Sprintf("Insertado actividad %s", id))
	return salida
}

func disconnectFromMongo(cliente *Client) {
	err := cliente.Disconnect(context.TODO())
	if err != nil {
		log.Error(err)
	}

	log.Info("Desconectado de mongo.")
}