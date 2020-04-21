package api

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"

	"github.com/mrojasb2000/fullstack/api/controllers"
	"github.com/mrojasb2000/fullstack/api/seed"
)

var server = controllers.Server{}

const http_port = "8080"
const http_server = "localhost"

func init() {
	// loads values from .env into the system
	if err := godotenv.Load(); err != nil {
		log.Print("sad .env file found")
	}
}

// Run - run application
func Run() {
	var err error
	err = godotenv.Load()
	if err != nil {
		log.Fatalf("Error getting env, %v", err)
	} else {
		fmt.Println("We are getting the env values")
	}

	server.Initialize(os.Getenv("DB_DRIVER"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_PORT"), os.Getenv("DB_HOST"), os.Getenv("DB_NAME"))

	seed.Load(server.DB)

	server.Run(http_port)
}
