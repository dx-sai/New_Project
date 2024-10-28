package main

import (
	"UserAPI/database"
	"UserAPI/kafka"
	"UserAPI/routes"
)

func main() {
	database.InitDB()

	kafka.Producer = kafka.InitProducer()
	defer kafka.Producer.Close()

	// Starting the consuming messages from Kafka topics
	go kafka.InitConsumer("signup")
	go kafka.InitConsumer("purchase")
	go kafka.InitConsumer("event_alerts")

	r := routes.SetupRouter()
	r.Run(":8080")
}
