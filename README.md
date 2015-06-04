# amqpsender
Command line to easily send messages to RabbitMQ

## Installation

### Latest source

```
$ go get -u github.com/rodrigosaito/amqpsender
```

### Package Managers (Ubuntu 14.04, CentOS 7)

Add my package cloud repository:

```
$ curl -s https://packagecloud.io/install/repositories/rodrigosaito/pkgs/script.rpm.sh | sudo bash
```

Then install the ```amqpsender``` package.

**Ubuntu 14.04 LTS:**

```
$ sudo apt-get install amqpsender
```

**CentOS 7:**

```
$ yum install amqpsender
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
