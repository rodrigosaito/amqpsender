package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/streadway/amqp"
)

const VERSION = "0.0.3"

func main() {
	var amqpUrl, exchange, routingKey, headers string
	var version bool
	flag.StringVar(&amqpUrl, "amqp-url", "amqp://guest:guest@localhost:5672", "rabbitmq url to send the messages to")
	flag.StringVar(&exchange, "exchange", "", "rabbitmq exchange to be used")
	flag.StringVar(&routingKey, "routing-key", "", "routing key to be used")
	flag.StringVar(&headers, "headers", "", "headers to be sent format: name1:value1,name2:value2")
	flag.BoolVar(&version, "v", false, "show current version")
	flag.Parse()

	args := flag.Args()

	if version {
		fmt.Println(VERSION)
		os.Exit(0)
	}

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
		if err = publish(channel, exchange, routingKey, parseHeaders(headers), args[0]); err != nil {
			panic(err)
		}
	} else {
		reader := bufio.NewReader(os.Stdin)

		for {
			line, err := reader.ReadString('\n')
			if err != nil {
				break
			}

			if err = publish(channel, exchange, routingKey, parseHeaders(headers), strings.TrimSpace(line)); err != nil {
				panic(err)
			}
		}
	}
}

func parseHeaders(h string) map[string]interface{} {
	headers := map[string]interface{}{}

	values := strings.Split(h, ",")
	for _, v := range values {
		name := strings.Split(v, ":")[0]
		value := strings.Split(v, ":")[1]
		headers[name] = value
	}

	return headers
}

func publish(channel *amqp.Channel, exchange, routingKey string, headers map[string]interface{}, message string) error {
	log.Printf("Publishing message: [%v]", message)
	err := channel.Publish(
		exchange,   // publish to an exchange
		routingKey, // routingKey
		false,      // mandatory
		false,      // immediate
		amqp.Publishing{
			Body:    []byte(message),
			Headers: headers,
		},
	)

	return err
}
