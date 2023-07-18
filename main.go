package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/eneassena/gointensivo-jul/internal/infra/database"
	"github.com/eneassena/gointensivo-jul/internal/usecase"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "db.sqlite3")
	if err != nil {
		log.Fatal(err)
	}
	orderRepository := database.NewOrderRepository(db)
	uc := usecase.NewCalculateFinalPrice(orderRepository)

	input := usecase.NewOrderInput("15", 50, 5)

	outPut, err := uc.Execute(*input)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(*outPut)
}
