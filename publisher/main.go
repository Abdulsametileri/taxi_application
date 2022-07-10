package main

import (
	"encoding/json"
	"fmt"
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

	if err = ch.ExchangeDeclare(
		"taxi-topic", // name
		"topic",      // type
		true,         // durable
		false,        // auto-deleted
		false,        // internal
		false,        // no-wait
		nil,          // arguments
	); err != nil {
		log.Panicln(err)
	}

	err = orderTaxi(ch, "taxi-topic", "taxi")
	err2 := orderTaxi(ch, "taxi-topic", "taxi.eco.#")
	if err != nil || err2 != nil {
		log.Panicln(err)
	}
}

func orderTaxi(ch *amqp.Channel, exchangeName, taxiID string) error {
	payload := fmt.Sprintf("example-message for %s", taxiID)
	messageID := uuid.NewString()

	jsonBytes, _ := json.Marshal(payload)

	err := ch.Publish(
		exchangeName, // exchange
		taxiID,       // routing key
		false,        // mandatory
		false,        // immediate
		amqp.Publishing{
			DeliveryMode: amqp.Persistent, // Specifies if the message should be persisted to disk or not.
			ContentType:  "application/json",
			Body:         jsonBytes,
			MessageId:    messageID, // message identifiers are an important aspect of traceability in messaging and distributed applications.,
		},
	)
	return err
}
