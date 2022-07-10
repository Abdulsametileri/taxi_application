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

	q1, _ := ch.QueueDeclare("taxi-inbox.1", true, true, false, false, nil)
	q2, _ := ch.QueueDeclare("taxi-inbox.2", true, true, false, false, nil)

	if err = ch.ExchangeDeclare("taxi-fanout", "fanout", true, false, false, false, nil); err != nil {
		log.Panicln(err)
	}

	ch.QueueBind(q1.Name, "", "taxi-fanout", false, nil)
	ch.QueueBind(q2.Name, "", "taxi-fanout", false, nil)

	ch.Publish("taxi-fanout", "", false, false, amqp.Publishing{
		ContentType:  "text/plain",
		DeliveryMode: amqp.Persistent,
		Body:         []byte("Hello everybody! This is an information message from the crew!"),
	})
}
