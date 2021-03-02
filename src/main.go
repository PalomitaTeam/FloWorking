package main



func main() {

	setUpLogs()

	collection := connectToMongo()

	getAllActivities(collection)
}