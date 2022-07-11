package main

import (
	"fmt"
	"github.com/Abdulsametileri/taxi-application/util"
	"github.com/google/uuid"
	amqp "github.com/rabbitmq/amqp091-go"
	"log"
	"strconv"
)

func main() {
	ch, cleanUp, err := util.CreateTCPConnectionAndOpenChannelOnIt()
	if err != nil {
		log.Panicln(err)
	}
	defer cleanUp()

	q, _ := ch.QueueDeclare("", false, false, true, false, nil)

	corrID := uuid.NewString()

	ch.Publish(
		"",          // exchange
		"rpc_queue", // routing key
		false,       // mandatory
		false,       // immediate
		amqp.Publishing{
			ContentType:   "text/plain",
			CorrelationId: corrID,
			ReplyTo:       q.Name,
			Body:          []byte("10"),
		})

	msgs, _ := ch.Consume(q.Name, "", true, false, false, false, nil)
	for d := range msgs {
		if corrID == d.CorrelationId {
			num, _ := strconv.Atoi(string(d.Body))
			fmt.Println(num)
		}
	}
}
