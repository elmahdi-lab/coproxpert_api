package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"ithumans.com/coproxpert/server"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		fmt.Printf("Failed to load .env file: %s\n", err)
	}

	err = server.ConnectDatabase()
	if err != nil {
		return
	}
	server.Start()

}
