package main
import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	log "github.com/sirupsen/logrus"
	"os"
	"time"
)

func connectToMongo() (*mongo.Client, *context.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	fmt.Println(os.Getenv("MONGO_URI"))
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(
		fmt.Sprintf(
			"mongodb://localhost:27017",
		),
	))

	if err != nil {
		log.Error("No se ha creado la conexión con Mongo" , err)
	}
	log.Info("Conectado a Mongo")
	return client, &ctx
}

func disconectFromMongo(client * mongo.Client) {
	err := client.Disconnect(context.TODO())

	if err != nil {
		log.Error("Desconectando Mongo", err)
	}
	log.Info("Connection to MongoDB closed.")
}


func insertActivity(col* mongo.Collection, ctx* context.Context, act Activity) {
	insertResult, err := col.InsertOne(context.TODO(), act)
	if err != nil {
		log.Error("Insertando Actividad", err)
	}
	fmt.Println("Inserted a single document: ", insertResult.InsertedID)
	/*var results []bson.M
	v, err := col.Find(*ctx, bson.M{})
	if err != nil {
		log.Println(err)
	}

	if err := v.All(*ctx, &results); err != nil {
		panic(err)
	}

	fmt.Println(v)*/
}

func getAllActivities(col* mongo.Collection) []*Activity {
	findOptions := options.Find()
	findOptions.SetLimit(2)

	var results []*Activity

	// Passing bson.D{{}} as the filter matches all documents in the collection
	cur, err := col.Find(context.TODO(), bson.D{{}}, findOptions)
	if err != nil {
		log.Error(err)
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
	log.SetFormatter(&log.JSONFormatter{})
	file, err := os.OpenFile("logs.log", os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		fmt.Println("No se ha podido crear el fichero de logs")
	}
	defer file.Close()
	log.SetOutput(file)
	db, ctx := connectToMongo()

	if ctx == nil {
		log.Println("ctx nil")
		fmt.Println("ctx nil")
	}
	fmt.Println("Conectado")

	collection := db.Database(os.Getenv("DB")).Collection(os.Getenv("COLLECTION"))
	log.Info(collection)

	act := Activity{
		name:        "prueba 1",
		duration:    10,
		description: "Esto es una descripcion de la prueba 1",
		status:       0,
		subActivity: newSubActivity("subActividad 1"),
	}

	insertActivity(collection, ctx, act);

	disconectFromMongo(db)
}
