package main

import (
	"github.com/Abdulsametileri/taxi-application/util"
	"log"
)

func main() {
	ch, cleanUp, err := util.CreateTCPConnectionAndOpenChannelOnIt()
	if err != nil {
		log.Panicln(err)
	}
	defer cleanUp()

	ch.Qos(1, 0, false)
	ch.QueueDeclare("", true, true, false, false, nil)
	ch.QueueBind("", "", "taxi-fanout", false, nil)

	msgs, err := ch.Consume("", "", false, false, false, false, nil)
	var forever chan struct{}

	go func() {
		for d := range msgs {
			log.Printf(" [x] %s", d.Body)
		}
	}()

	log.Printf(" [*] Waiting for logs. To exit press CTRL+C")
	<-forever
}
