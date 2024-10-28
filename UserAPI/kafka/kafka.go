package kafka

import (
	"log"
	"os"

	"github.com/Shopify/sarama"
	"github.com/joho/godotenv"
)

var Producer sarama.SyncProducer

func InitProducer() sarama.SyncProducer {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error loading the .env file")
	}

	broker := os.Getenv("KAFKA_BROKER")
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true

	producer, err := sarama.NewSyncProducer([]string{broker}, config)
	if err != nil {
		log.Fatal("Error creatin the kafka producer:", err)
	}
	return producer
}

func SendMessage(topic, message string) {
	msg := &sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.StringEncoder(message),
	}

	_, _, err := Producer.SendMessage(msg)
	if err != nil {
		log.Println("error sending kafka message", err)
	}
}
func InitConsumer(topic string) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	broker := os.Getenv("KAFKA_BROKER")
	config := sarama.NewConfig()
	config.Consumer.Return.Errors = true

	consumer, err := sarama.NewConsumer([]string{broker}, config)
	if err != nil {
		log.Fatal("Error creating Kafka consumer:", err)
	}
	defer consumer.Close()

	partitionConsumer, err := consumer.ConsumePartition(topic, 0, sarama.OffsetNewest)
	if err != nil {
		log.Fatal("Error creating partition consumer:", err)
	}
	defer partitionConsumer.Close()

	log.Println("Consumer started for topic:", topic)

	// Consume the messages
	for message := range partitionConsumer.Messages() {
		log.Printf("Message received from topic %s: %s", topic, string(message.Value))
	}
}
