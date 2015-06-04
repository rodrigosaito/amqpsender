# amqpsender
Command line for easily send messages to RabbitMQ

## Installation

To install from the latest source:

```
$ go get -u github.com/rodrigosaito/amqpsender
```

## Usage

### Sending a single message

To send a single message to a specific queue:

```
$ amqpsender -routing-key some_queue some_message
```

If you want to use a different exchange other than the default:

```
$ amqpsender -exchange another_exchange -routing_key some_routing_key some_message
```

If you don't want to use default amqp connection string:

```
$ amqpsender -amqp-url amqp://user:password@host:5672 -exchange another_exchange -routing_key some_routing_key some_message
```

### Sending multiple messages

You can also pipe the content of a file to ```amqpsender```:

```
$ cat file | amqpsender -routing-key some_routing_key
```

Then every single line will be sent as a single message to rabbitmq
