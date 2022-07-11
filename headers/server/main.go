package main

import (
	"github.com/Abdulsametileri/taxi-application/util"
	amqp "github.com/rabbitmq/amqp091-go"
	"log"
)

func main() {
	ch, cleanUp, err := util.CreateTCPConnectionAndOpenChannelOnIt()
	if err != nil {
		log.Panicln(err)
	}
	defer cleanUp()

	payload := `{ "latitude": 0.0, "longitude": 3.0 }`

	ch.Publish(
		"taxi_header_exchange",
		"",
		false,
		false,
		amqp.Publishing{
			Headers: amqp.Table{
				"version": "0.1b",
				"system":  "taxi",
			},
			ContentType:  "application/json",
			DeliveryMode: amqp.Persistent,
			Body:         []byte(payload),
		},
	)
}
