package main



func main() {

	f := setUpLogs()

	defer closeLogFile(f)

	cliente, collection := connectToMongo()

	getAllActivities(collection)

	disconnectFromMongo(cliente)
}