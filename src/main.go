package main

func main() {

	f := setUpLogs()

	defer closeLogFile(f)


	cliente, collection := connectToMongo()

	a := Activity{
		id:          nil,
		name:        "prueba",
		duration:    90,
		description: "desde el main",
		status:      0,
	}

	getAllActivities(collection)

	insertActivity(collection, a)

	getAllActivities(collection)

	disconnectFromMongo(cliente)
}