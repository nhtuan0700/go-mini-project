package main

import (
	"myproject/app/api"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load("./.env.local"); err != nil {
		panic(err)
	}
	api.RunServer()
}
