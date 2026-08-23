package main

import (
	"context"
	"log"

	"github.com/levisantosp/altamira-participa/api/db"
	"github.com/levisantosp/altamira-participa/api/utils"
)

func main() {
	utils.LoadEnv()
	db.Connect()

	err := db.Client.Schema.Create(context.Background())
	if err != nil {
		log.Fatal(err)
	}
}
