package main

import (
	"github.com/Abdulsametileri/taxi-application/util"
	"github.com/google/uuid"
	amqp "github.com/rabbitmq/amqp091-go"
	"log"
)

func main() {
	ch, cleanUp, err := util.CreateTCPConnectionAndOpenChannelOnIt()
	if err != nil {
		log.Panicln(err)
	}
	defer cleanUp()

	taxiID := "taxi.1"
	q, err := ch.QueueDeclare(
		taxiID, // name
		true,   // durable
		false,  // delete when unused
		false,  // exclusive
		false,  // no-wait
		nil,    // arguments
	)
	if err != nil {
		log.Panicln("Failed to declare a queue")
	}

	if err = orderTaxi(ch, q.Name); err != nil {
		log.Panicln(err)
	}
}

func orderTaxi(ch *amqp.Channel, queueName string) error {
	payload := "example-message"
	messageID := uuid.NewString()

	err := ch.Publish(
		"",        // exchange
		queueName, // routing key
		false,     // mandatory
		false,     // immediate
		amqp.Publishing{
			DeliveryMode: amqp.Persistent, // Specifies if the message should be persisted to disk or not.
			ContentType:  "text/plain",
			Body:         []byte(payload),
			MessageId:    messageID, // message identifiers are an important aspect of traceability in messaging and distributed applications.,
		},
	)
	return err
}
