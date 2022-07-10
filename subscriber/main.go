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

	q, err := ch.QueueDeclare("taxi", true, true, false, false, nil)
	qEco, err2 := ch.QueueDeclare("taxi.eco", true, true, false, false, nil)
	if err != nil || err2 != nil {
		log.Panicln(err)
	}

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

	err = ch.QueueBind(q.Name, "taxi", "taxi-topic", false, nil)
	err2 = ch.QueueBind(qEco.Name, "taxi.eco.#", "taxi-topic", false, nil)
	if err != nil || err2 != nil {
		log.Panicln(err)
	}

	var forever chan struct{}

	go consumeFrom(q.Name, ch)
	go consumeFrom(qEco.Name, ch)

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}

func consumeFrom(queueName string, ch *amqp.Channel) {
	msgs, _ := ch.Consume(
		queueName, // queue
		"",        // consumer
		true,      // auto-ack
		false,     // exclusive
		false,     // no-local
		false,     // no-wait
		nil,       // args
	)

	for d := range msgs {
		log.Printf("Received a message: %s", d.Body)
	}
}
