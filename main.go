package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"ithumans.com/coproxpert/cmd"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		fmt.Printf("Failed to load .env file: %s\n", err)
	}

	err = cmd.ConnectDatabase()
	if err != nil {
		return
	}
	cmd.Start()

}
