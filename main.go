package main

import (
	"bufio"
	"flag"
	"log"
	"os"
	"strings"

	"github.com/streadway/amqp"
)

func main() {
	var amqpUrl, exchange, routingKey string
	flag.StringVar(&amqpUrl, "amqp-url", "amqp://guest:geust@localhost:5672", "rabbitmq url to send the messages to")
	flag.StringVar(&exchange, "exchange", "", "rabbitmq exchange to be used")
	flag.StringVar(&routingKey, "routing-key", "", "routing key to be used")
	flag.Parse()

	args := flag.Args()

	conn, err := amqp.Dial(amqpUrl)
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	channel, err := conn.Channel()
	if err != nil {
		panic(err)
	}
	defer channel.Close()

	log.Printf("Using amqp-url: [%v], exchange: [%v], routing-key: [%v]", amqpUrl, exchange, routingKey)

	if len(args) > 0 {
		if err = publish(channel, exchange, routingKey, args[0]); err != nil {
			panic(err)
		}
	} else {
		reader := bufio.NewReader(os.Stdin)

		for {
			line, err := reader.ReadString('\n')
			if err != nil {
				break
			}

			if err = publish(channel, exchange, routingKey, strings.TrimSpace(line)); err != nil {
				panic(err)
			}
		}
	}
}

func publish(channel *amqp.Channel, exchange, routingKey, message string) error {
	log.Printf("Publishing message: [%v]", message)
	err := channel.Publish(
		exchange,   // publish to an exchange
		routingKey, // routingKey
		false,      // mandatory
		false,      // immediate
		amqp.Publishing{
			Body: []byte(message),
		},
	)

	return err
}
