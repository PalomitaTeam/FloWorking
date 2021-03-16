package main



func main() {

	f := setUpLogs()

	defer closeLogFile(f)

	cliente, collection, _ := connectToMongo()

	getAllActivities(collection)

	disconnectFromMongo(cliente)
}