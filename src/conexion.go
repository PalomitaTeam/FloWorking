package main
import (
	"context"
	"fmt"
	"log"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"os"
	"time"
)

const(
	DB = "FloWorking"
	COLLECTION = "Activities"
)

func connectToMogo() (*mongo.Client) { //Cambiar <password> por la contraseña y myFirstDatabase por db

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	fmt.Println(os.Getenv("MONGO_URI"))
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(
		"mongodb+srv:// FloWork:FloWorkingAngelJJ@cluster0.lujjr.mongodb.net/myFirstDatabase?retryWrites=true&w=majority",
	))

	if err != nil {
		log.Fatal(err)
	}

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
}

func main() {
	db := connectToMogo()

	fmt.Println("Conectado")
	collection := db.Database(DB).Collection(COLLECTION)

	results := getAllActivities(collection)

	fmt.Println(results)

	disconectFromMongo(db)
}