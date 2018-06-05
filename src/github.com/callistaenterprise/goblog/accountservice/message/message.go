package message

import (
	"github.com/streadway/amqp"
	"fmt"
	"log"
)

type IMessagingClient interface {
	ConnectToBroker(connectionString string)
	// Publish(msg []byte, exchangeName string, exchangeType string) error
	PublishOnQueue(msg []byte, queueName string) error
	// Subscribe(exchangeName string, exchangeType string, consumerName string, handlerFunc func(amqp.Delivery)) error
	SubscribeToQueue(queueName string, consumerName string, handerFunc func(amqp.Delivery)) error
	// Close()
}


type MessagingClient struct {
	conn *amqp.Connection
}

func(m *MessagingClient) ConnectToBroker(connectionString string) {
	if connectionString == "" {
		panic("Cannot initialize connection to broker, connectionString not set. Have you initialized?")
	}

	var err error
	m.conn, err = amqp.Dial(fmt.Sprintf("%s/", connectionString))
	if err != nil {
		panic(" connect to AMQP failed at" + connectionString)
	}
}

func (m *MessagingClient) PublishOnQueue(body []byte, queueName string) error {
	if m.conn == nil {
		panic("Tried to send message before connection was initialized. Don't do that.")
	}

	ch, err := m.conn.Channel()
	defer ch.Close()

	queue, err := ch.QueueDeclare(
		queueName,
		false,
		false,
		false,
		false,
		nil,
	)

	err = ch.Publish(
		"",
		queue.Name,
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body: body,
		})

	fmt.Printf("A message was sent to queue %v: %v", queueName, body)

	return err
}

func(m *MessagingClient) SubscribeToQueue(queueName string, consumerName string, handlerFunc func(amqp.Delivery)) error {
	ch, err := m.conn.Channel()
	if err != nil {
		panic("Failed to open channel")
	}

	log.Printf("Declaring Queue(%s)", queueName)

	queue, err := ch.QueueDeclare(
		queueName,
		false,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		panic("Failed to register an Queue")
	}

	msgs, err := ch.Consume(
		queue.Name,
		consumerName,
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		panic("Failed to register a consumer")
	}

	go consumerLoop(msgs, handlerFunc)

	return nil
}

func consumerLoop(deliveries <-chan amqp.Delivery, handlerFunc func(d amqp.Delivery)) {
	for d := range deliveries {
		handlerFunc(d)
	}
}