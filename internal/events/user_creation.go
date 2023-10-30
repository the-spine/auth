package events

import (
	"auth/internal/config"
	"context"
	"fmt"
	"log"
	"net"
	"strconv"
	"time"

	"github.com/segmentio/kafka-go"
)

const (
	userCreationTopic = "user_creation"
	partition         = 0
	commitInterval    = time.Second
)

func CreateTopicIfNotExists(config *config.Config) error {
	network := "tcp"
	address := fmt.Sprintf("%s:%d", config.Kafka.Host, config.Kafka.Port)

	conn, err := kafka.Dial(network, address)
	if err != nil {
		return err
	}
	defer conn.Close()

	controller, err := conn.Controller()
	if err != nil {
		return err
	}
	var controllerConn *kafka.Conn
	controllerConn, err = kafka.Dial("tcp", net.JoinHostPort(controller.Host, strconv.Itoa(controller.Port)))
	if err != nil {
		return err
	}
	defer controllerConn.Close()

	// Check if the topic already exists
	topics, err := controllerConn.ReadPartitions()
	if err != nil {
		return err
	}

	for _, topic := range topics {
		if topic.Topic == userCreationTopic {
			// The topic already exists, no need to create it
			return nil
		}
	}

	topicConfigs := []kafka.TopicConfig{
		{
			Topic:             userCreationTopic,
			NumPartitions:     1,
			ReplicationFactor: 1,
		},
	}

	err = controllerConn.CreateTopics(topicConfigs...)
	if err != nil {
		return err
	}

	return nil
}

func ConsumeUserCreationTopic(config *config.Config) {

	address := fmt.Sprintf("%s:%d", config.Kafka.Host, config.Kafka.Port)

	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{address},

		Topic:    userCreationTopic,
		MinBytes: 10,   // 10 bytes minimum
		MaxBytes: 10e6, // 10MB maximum
	})

	for {
		m, err := r.FetchMessage(context.Background())
		if err != nil {
			log.Printf("Error fetching message: %v", err)
			break
		}
		// create user here

		r.CommitMessages(context.Background(), m)
	}

	r.Close()
}
