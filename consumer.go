package main

import (
	"fmt"

	"github.com/streadway/amqp"
)

// import (
// 	"fmt"

// 	"github.com/streadway/amqp"
// )

func main() {
	fmt.Println("consumer Application")

	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	defer ch.Close()

	msgs, err := ch.Consume(
		"TestQueue",
		"",
		true,
		false,
		false,
		false,
		nil,
	)

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			fmt.Printf("Recieved Messages:  %s\n", d.Body)
		}
	}()

	fmt.Println("successfully consumed to our rabbitMQ instance")
	fmt.Println(" [*] - waithing for messages")

	<-forever

}
