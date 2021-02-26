package main
import (
	"context"
	"fmt"
	"log"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func connectToMogo() (*mongo.Client) {						//Cambiar <password> por la contraseña y myFirstDatabase por db
	clientOptions := options.Client().ApplyURI("server de mongo",)
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

func getAllActivities(col* mongo.Collection) []*Activity {
	findOptions := options.Find()
	findOptions.SetLimit(2)

	var results []*Activity

	// Passing bson.D{{}} as the filter matches all documents in the collection
	cur, err := col.Find(context.TODO(), bson.D{{}}, findOptions)
	if err != nil {
		log.Fatal(err)
	}

	// Finding multiple documents returns a cursor
	// Iterating through the cursor allows us to decode documents one at a time
	for cur.Next(context.TODO()) {

		// create a value into which the single document can be decoded
		var elem Activity
		err := cur.Decode(&elem)
		if err != nil {
			log.Fatal(err)
		}

		results = append(results, &elem)
	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	return results
	/*
	var results []interface{}
	err := col.FindOne(context.TODO(),"{duration:90}").Decode(&results)
	if err != nil {
		log.Fatal(err)
	}
	return results
	*/

}

func main() {
	db := connectToMogo()
	collection := db.Database("FloWorking").Collection("Activities")

	/*act := Activity{
		id:          "a",
		name:        "prueba 1",
		duration:    10,
		description: "Esto es una descripcion de la prueba 1",
		status:       pending,
		subActivity: newSubActivity("subActividad 1"),
	}*/

	//(collection, act)
	fmt.Println("Insertada la actividad")
	var results []*Activity
	results= getAllActivities(collection)
	fmt.Printf("Found multiple documents (array of pointers): %+v\n", results)
	disconectFromMongo(db)
}