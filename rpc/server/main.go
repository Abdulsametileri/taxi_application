package main

import (
	"github.com/Abdulsametileri/taxi-application/util"
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

	q, _ := ch.QueueDeclare("rpc_queue", false, false, false, false, nil)
	msgs, _ := ch.Consume(q.Name, "", false, false, false, false, nil)

	var forever chan struct{}
	go func() {
		for msg := range msgs {
			n, _ := strconv.Atoi(string(msg.Body))

			n = n * n

			ch.Publish("", msg.ReplyTo, false, false, amqp.Publishing{
				ContentType:   "text/plain",
				CorrelationId: msg.CorrelationId,
				Body:          []byte(strconv.Itoa(n)),
			})
		}
	}()
	<-forever
}
