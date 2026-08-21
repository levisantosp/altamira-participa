package db

import (
	"database/sql"
	"log"

	"entgo.io/ent/dialect"
	entsql "entgo.io/ent/dialect/sql"
	"github.com/levisantosp/altamira-participa/ent/generated"
	"github.com/levisantosp/altamira-participa/utils"

	_ "github.com/jackc/pgx/v5/stdlib"
)

var Client *generated.Client

func Connect() {
	db, err := sql.Open("pgx", utils.Env.DatabaseURL)
	if err != nil {
		log.Fatal(err)
	}

	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	Client = generated.NewClient(
		generated.Driver(entsql.OpenDB(dialect.Postgres, db)),
	)
}
