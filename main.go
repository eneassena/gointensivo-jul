package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"

	"github.com/eneassena/gointensivo-jul/internal/infra/database"
	"github.com/eneassena/gointensivo-jul/internal/usecase"
	"github.com/eneassena/gointensivo-jul/pkg/rabbitmq"
	_ "github.com/mattn/go-sqlite3"

	amqp "github.com/rabbitmq/amqp091-go"
)

func main() {
	db, err := sql.Open("sqlite3", "db.sqlite3")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	orderRepository := database.NewOrderRepository(db)
	uc := usecase.NewCalculateFinalPrice(orderRepository)

	// criar a conexão com o rabbitmq
	ch, err := rabbitmq.OpenChannel()
	if err != nil {
		panic(err)
	}
	defer ch.Close()

	messengerRabbitmqChannel := make(chan amqp.Delivery)
	go rabbitmq.Consume(ch, messengerRabbitmqChannel) // executando a fila // trava // T2

	rabbitmqWorker(messengerRabbitmqChannel, uc) // T1
}

func rabbitmqWorker(msgChan chan amqp.Delivery, uc *usecase.CalculateFinalPrice) {
	fmt.Println("starting rabbitmq")
	for msg := range msgChan {
		var input usecase.OrderInput
		err := json.Unmarshal(msg.Body, &input)
		if err != nil {

			panic(err)
		}
		output, err := uc.Execute(input)
		if err != nil {
			panic(err)
		}
		msg.Ack(false)
		fmt.Println("Mensagem processada e salva no banco:", output)
	}
}
