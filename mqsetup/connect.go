package mqsetup

import (
	"fmt"
	"os"

	//amqp "github.com/rabbitmq/amqp091-go"
	"github.com/streadway/amqp"
	//"context"
	"log"
)

func ErrInConnect(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}
func Read() []byte {
	file, err := os.ReadFile("scrap.txt")
	if err != nil {
		log.Println(err)
	}
	return file
}

func MqSetup() {

	conn, err := amqp.Dial("amqp://guest@localhost:5672/")
	ErrInConnect(err, "failed to connect rabbit mq")
	defer func(conn *amqp.Connection) {
		err := conn.Close()
		if err != nil {
			panic(err)
		} else {
			fmt.Println("connected")
		}
	}(conn)

	ch, err := conn.Channel()

	ErrInConnect(err, "failed to open a channel")
	defer func(ch *amqp.Channel) {
		err := ch.Close()
		if err != nil {
			println(err)
			panic(err)
		}
	}(ch)
	q, err := ch.QueueDeclare(
		"testque",
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	fmt.Println(q)

	err = ch.Publish(
		"",
		"testque",
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        Read(),
		},
	)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	fmt.Println("success")
}
