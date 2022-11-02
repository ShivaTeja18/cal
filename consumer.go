package main

import (
	"cal/mqsetup"
	"encoding/json"
	"fmt"
	"github.com/streadway/amqp"
	"log"
	"os"
)

//func file() *os.File {
//	scrap, err := os.Create("scrap.txt")
//	if err != nil {
//		log.Println(err)
//	}
//	err = os.WriteFile("scrap.txt")
//	if err != nil {
//		return nil
//	}
//	return scrap
//}

func Consumer() {
	fmt.Println("Consumer Application")
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	ch, err := conn.Channel()
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	mqsetup.ErrInConnect(err, "failed to open channel")
	defer func(ch *amqp.Channel) {
		err := ch.Close()
		if err != nil {
			fmt.Println(err)
			panic(err)
		}
	}(ch)
	msgs, err := ch.Consume(
		"testque",
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	limitless := make(chan bool)
	go func() {
		var un interface{}
		for m := range msgs {
			err = json.Unmarshal(m.Body, &un)
			if err != nil {
				log.Println(err)
			}
			fmt.Println(un)
			a := un.(map[string]interface{})
			for k, v := range a {
				switch val := v.(type) {
				case string:
					fmt.Println(k, "is string", val)
				case float64:
					fmt.Printf("float %v %d\n", k, int(val))
				default:
					fmt.Println("unknown", k, v)
				}
			}
		}
	}()
	fmt.Println("success")
	fmt.Println("[*] - waiting for msgs")
	<-limitless
}

func main() {
	fmt.Println(os.Args)
	Consumer()
}
