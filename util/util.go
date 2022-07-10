package util

import (
	"errors"
	amqp "github.com/rabbitmq/amqp091-go"
)

func CreateTCPConnectionAndOpenChannelOnIt() (*amqp.Channel, func(), error) {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672")
	if err != nil {
		return nil, nil, errors.New("failed to connect RabbitMQ")
	}
	cleanUp := func() {
		conn.Close()
	}

	ch, err := conn.Channel()
	if err != nil {
		return nil, cleanUp, errors.New("failed to open a channel")
	}
	cleanUp = func() {
		conn.Close()
		ch.Close()
	}

	return ch, cleanUp, nil
}
