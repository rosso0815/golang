package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	//	"sync"
	//	"time"
	//	"encoding/base64"
	"github.com/codegangsta/cli"
	//	"github.com/streadway/amqp"
	//	"github.com/go-stomp/stomp"
)

//var log *log.Logger

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
		panic(fmt.Sprintf("%s: %s", msg, err))
	}
}

//func prepare(amqpURI string, queueName string, createQueue bool) {
//conn, err := amqp.Dial(amqpURI)
//	conn, err := stomp.Dial("tcp", "localhost:61613")
//	failOnError(err, "Failed to connect to AMQP broker")

//	channel, err := conn.Channel()
//	failOnError(err, "Failed to open a channel")

//	if createQueue == true {
//		_, err = channel.QueueDeclare(
//			queueName, // name
//			true,      // durable
//			false,     // delete when unused
//			false,     // exclusive
//			false,     // no-wait
//			nil,       // arguments
//		)
//		failOnError(err, "Failed to declare queue")
//	}

//	return conn, channel
//}

func publish(c *cli.Context) {
	log.Print("@@@ publish")

	queueName := c.Args().First()
	if queueName == "" {
		fmt.Println("Please provide name of the queue")
		os.Exit(1)
	}
	log.Print("use queue " + queueName)

	//	conn, channel := prepare(c.String("amqpuri"), queueName, !c.Bool("no-create-queue"))

	//	defer conn.Close()
	//	defer channel.Close()

	// It is better to have a durable delivery mode and let user disable it
	// even though it is not the RabbitMQ default
	//	deliveryMode := uint8(2)
	///	if c.Bool("transient") {
	//		deliveryMode = 1
	//	}

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		log.Print("got line " + line)
		//		msgBody := []byte(line)
		//		if c.Bool("base64") {
		//			var encodeErr error
		//			msgBody, encodeErr = base64.StdEncoding.DecodeString(line)
		//			failOnError(encodeErr, "Fail to decode base64")
		//		}

		//		err := channel.Publish(
		//			c.String("exchange"), // exchange
		//			queueName,            // routing key
		//			false,                // mandatory
		//			false,                // immediate
		//			amqp.Publishing{
		//				ContentType:  "text/plain",
		//				Body:         msgBody,
		//				DeliveryMode: deliveryMode,
		//			})
		//
		//		failOnError(err, "Failed to publish a message")
		//		fmt.Println(line)
	}
	err := scanner.Err()
	failOnError(err, "Failed to read from stdin")
}

func consume(c *cli.Context) {
	queueName := c.Args().First()
	if queueName == "" {
		fmt.Println("Please provide name of the queue")
		os.Exit(1)
	}

	//	conn, channel := prepare(c.String("amqpuri"), queueName, !c.Bool("no-create-queue"))
	//	defer conn.Close()
	//	defer channel.Close()

	//	var mutex sync.Mutex
	//	unackedMessages := make([]amqp.Delivery, 100)

	//	msgs, err := channel.Consume(
	//		queueName,         // queue
	//		"",                // consumer
	//		c.Bool("autoack"), // auto-ack
	//		false,             // exclusive
	//		false,             // no-local
	//		false,             // no-wait
	//		nil,               // args
	//	)
	//	failOnError(err, "Failed to register consumer")

	//	ackMessages := func() {
	//		scanner := bufio.NewScanner(os.Stdin)
	//		for scanner.Scan() {
	//ackedLine := scanner.Text()

	// O(n²) complexity for the win!
	//			mutex.Lock() // use channels some day
	//			for i, msg := range unackedMessages {
	//				unackedLine := fmt.Sprintf("%s", msg.Body)
	//				if c.Bool("base64") {
	//					unackedLine = base64.StdEncoding.EncodeToString([]byte(unackedLine))
	//				}
	//				if unackedLine == ackedLine {
	//					msg.Ack(false)
	//
	//					// discard message
	//					unackedMessages = append(unackedMessages[:i], unackedMessages[i+1:]...)
	//					break
	//				}
	//			}
	//			mutex.Unlock()

	//		}
	//		err := scanner.Err()
	//		failOnError(err, "Failed to read from stdin")
	//	}

	//forever := make(chan bool)

	//	consumeMessages := func() {
	//		timeout := time.Second * time.Duration(c.Int("timeout"))
	//		for {
	//			select {
	//			case msg := <-msgs:
	//				if !c.Bool("autoack") {
	//					mutex.Lock()
	//					unackedMessages = append(unackedMessages, msg)
	//					mutex.Unlock()
	//				}
	//				line := fmt.Sprintf("%s", msg.Body)
	//				if c.Bool("base64") {
	//					line = base64.StdEncoding.EncodeToString([]byte(line))
	//				}
	//				fmt.Println(line)
	//			case <-time.After(timeout):
	//				if c.Bool("non-blocking") {
	//					forever <- false
	//					return
	//				}
	//			}
	//		}
	//	}

	//	if c.Bool("autoack") {
	///		go consumeMessages()
	//	} else {
	//		go ackMessages()
	//		go consumeMessages()
	//	}
	//	<-forever
}

func main() {
	app := cli.NewApp()
	app.Name = "pipestomp"
	app.Usage = "forward unix pipe to stomp broker"
	app.Version = "0.0.1"

	globalFlags := []cli.Flag{
	//		cli.StringFlag{
	//			Name:   "amqpuri",
	//			Value:  "amqp://guest:guest@localhost:5672/",
	//			Usage:  "AMQP URI",
	//			EnvVar: "AMQP_URI",
	//		},
	//		cli.StringFlag{
	//			Name:   "exchange",
	//			Value:  "",
	//			Usage:  "AMQP Exchange to publish to (default: \"\")",
	//			EnvVar: "AMQP_EXCHANGE",
	//		},
	//		cli.BoolFlag{
	//			Name:  "no-create-queue",
	//			Usage: "Don't create queue",
	//		},
	//		cli.BoolFlag{
	//			Name:  "autoack",
	//			Usage: "Ack all received messages directly",
	//		},
	//		cli.BoolFlag{
	//			Name:  "non-blocking",
	//			Usage: "Stop consumer after timeout",
	//		},
	//		cli.BoolFlag{
	//			Name:  "transient",
	//			Usage: "Publish messages with transient delivery mode",
	//		},
	//		cli.IntFlag{
	//			Name:  "timeout",
	//			Value: 1,
	//			Usage: "Timeout to wait for messages",
	//		},
	//		cli.BoolFlag{
	//			Name:  "base64",
	//			Usage: "Encode to Base64 string in consumer mode. Decode from Base64 string in publish mode",
	//		},
	}

	app.Commands = []cli.Command{
		{
			Name:    "publish",
			Aliases: []string{"p"},
			Usage:   "Publish messages to queue",
			Flags:   globalFlags,
			Action:  publish,
		},
		//		{
		//			Name:    "consume",
		//			Flags:   globalFlags,
		//			Aliases: []string{"c"},
		//			Usage:   "Consume messages from queue",
		//			Action:  consume,
		//		},
	}

	app.Run(os.Args)

	log.Print("done")
}
