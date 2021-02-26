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

func connectToMongo() (*mongo.Client, *context.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	fmt.Println(os.Getenv("MONGO_URI"))
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(
		fmt.Sprintf(
			"MONGO_URI=mongodb+srv://%s:%s@cluster0.lujjr.mongodb.net/myFirstDatabase?retryWrites=true&w=majority",
			os.Getenv("USER"), os.Getenv("PASS"),
		),
	))

	if err != nil {
		log.Fatal(err)
	}

	return client, &ctx
}

func disconectFromMongo(client * mongo.Client) {
	err := client.Disconnect(context.TODO())

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connection to MongoDB closed.")
}


func insertActivity(col* mongo.Collection, ctx* context.Context, act Activity) {
	var results []bson.M
	v, err := col.Find(*ctx, bson.M{})
	if err != nil {
		log.Println(err)
	}

	if err := v.All(*ctx, &results); err != nil {
		panic(err)
	}

	fmt.Println(v)
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
	db,ctx := connectToMongo()

	fmt.Println("Conectado")

	collection := db.Database(COLLECTION).Collection(DB)

	act := Activity{
		id:          "78",
		name:        "prueba 1",
		duration:    10,
		description: "Esto es una descripcion de la prueba 1",
		status:       0,
		subActivity: newSubActivity("subActividad 1"),
	}

	insertActivity(collection, ctx, act);

	disconectFromMongo(db)
}
