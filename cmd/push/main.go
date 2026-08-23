package main

import (
	"context"
	"log"

	"github.com/levisantosp/altamira-participa/db"
	"github.com/levisantosp/altamira-participa/utils"
)

func main() {
	utils.LoadEnv()
	db.Connect()

	err := db.Client.Schema.Create(context.Background())
	if err != nil {
		log.Fatal(err)
	}
}
