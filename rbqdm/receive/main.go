package main

import (
	"fmt"
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

func main() {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/xx")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"hello", // name
		false,   // durable
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)
	failOnError(err, "Failed to declare a queue")
	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		false,  // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	failOnError(err, "Failed to register a consumer")

	var forever chan struct{}

	msgHandle := func() {
		defer func() {
			if err := recover(); err != nil {
				fmt.Printf("panic is %v", err)
			}
		}()

		for d := range msgs {
			handleMsg(d, handleMsgFunc(func(delivery amqp.Delivery) error {
				log.Println("func1...")
				return nil
			}), handleMsgFunc(func(delivery amqp.Delivery) error {
				log.Println("func1...")
				return nil
			}))
		}
	}
	go func() {
		msgHandle()
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}

type handleMsgFunc func(delivery amqp.Delivery) error

func (h handleMsgFunc) handle(delivery amqp.Delivery) {
	log.Println("start")
	err := h(delivery)
	if err != nil {
		panic(fmt.Sprintf("msg error %s", err.Error()))
	}
	log.Println("end")
}

func handleMsg(d amqp.Delivery, hmf ...handleMsgFunc) {
	for _, h := range hmf {
		h.handle(d)
	}
	err := d.Ack(true)
	if err != nil {
		panic(fmt.Sprintf("msg error %s", err.Error()))
	}
}
