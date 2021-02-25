package main
import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func connectToMogo() (*mongo.Client) {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
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

func insertActivity(col* Collection, act Activity) {
	insertResult, err := col.InsertOne(context.TODO(), act)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted a single document: ", insertResult.InsertedID)
}

func getAllActivities(col* Collection) []Activity {
	var results []Activity
	err := col.Find(nil).All(&results)
	if err != nil {
		log.Fatal(err)
	}
	return results
}

func main() {
	db := connectToMogo()

	collection := db.Database("floworking").Collection("activities")

	act := Activity{
		id:          nil,
		name:        "prueba 1",
		duration:    10,
		description: "Esto es una descripcion de la prueba 1",
		state:       false,
		subActivity: newSubActivity("subActividad 1"),
	}

	insertActivity(collection, act)
	fmt.Println("Insertada la actividad")
	getAllActivities(collection)
	disconectFromMongo(db)
}