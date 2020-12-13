package main

import (
	"../lambda"
	"../api"
	"../database"

	"os"
	"log"
)

func init() {
	if os.Getenv("USE_AWS") == "true" {
		lambda.Client = lambda.AWSImplementation{}
	} else {
		lambda.Client = lambda.MockAWS{}
	}

	if os.Getenv("USE_DB") == "true" {
		database.SetDatabase(database.DBImplementation{})
	} else {
		database.SetDatabase(database.DBMock{})
	}
}


func main() {

	if err := api.StartServer(); err != nil {
		log.Fatal("Failed to start server: ", err)
	}
	
}


