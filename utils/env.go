package utils

import (
	"log"

	"github.com/Oudwins/zog"
	"github.com/Oudwins/zog/zenv"
	"github.com/joho/godotenv"
)

type TEnv struct {
	DatabaseURL string `zog:"DATABASE_URL"`
	RedisAddr string `zog:"REDIS_ADDR"`
}

var Env TEnv

func LoadEnv() {
	godotenv.Load()

	schema := zog.Struct(zog.Shape{
		"DatabaseURL": zog.String().URL().Required(),
		"RedisAddr": zog.String().URL().Required(),
	})

	err := schema.Parse(zenv.NewDataProvider(), Env)
	if err != nil {
		log.Fatal(err)
	}
}
