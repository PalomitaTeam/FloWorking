package main
import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func connectToMogo() (*mongo.Client) {						//Cambiar <password> por la contraseña
	clientOptions := options.Client().ApplyURI("server de mongo")
	client, err := mongo.Connect(context.TODO(), clientOptions)

	// Los errores habría que sacarlos de aqui (?)
	if err != nil {
		log.Fatal(err)
	}

	// Check the connection
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")
	return client
}

func disconectFromMongo(client * mongo.Client) {
	err := client.Disconnect(context.TODO())

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connection to MongoDB closed.")
}


func insertActivity(col* mongo.Collection, act Activity) {
	insertResult, err := col.InsertOne(context.TODO(), act)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted a single document: ", insertResult.InsertedID)
}

func getAllActivities(col* mongo.Collection) []interface{} {
	var results []interface{}
	err := col.FindOne(context.TODO(),"{duration:90}").Decode(&results)
	if err != nil {
		log.Fatal(err)
	}
	return results
}

func main() {
	db := connectToMogo()
	collection := db.Database("FloWorking").Collection("Activities")

	act := Activity{
		id:          "a",
		name:        "prueba 1",
		duration:    10,
		description: "Esto es una descripcion de la prueba 1",
		status:       pending,
		subActivity: newSubActivity("subActividad 1"),
	}

	insertActivity(collection, act)
	fmt.Println("Insertada la actividad")
	getAllActivities(collection)
	disconectFromMongo(db)
}