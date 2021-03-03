package main

import log "github.com/sirupsen/logrus"

func main() {

	f := setUpLogs()

	defer closeLogFile(f)

	log.Info("hola esto e sun indo")

	//cliente, collection := connectToMongo()

	//getAllActivities(collection)

	//disconnectFromMongo(cliente)
}