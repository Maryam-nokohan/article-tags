package configs

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func Newconfig() (*Config, error) {
	err := godotenv.Load(".env")
	if err != nil {
		log.Println("No .env file found, using system environment variables")
	}

	dbName := os.Getenv("DB_NAME")
	gRPCPort:= os.Getenv("GRPC_PORT")
	uri := os.Getenv("URI")

	if uri == "" || dbName == "" || gRPCPort == "" {
		log.Fatal("setup your .env")
	}

	return &Config{
		DBName: dbName,
		GRPC_Port: gRPCPort,
		URI:    uri,
	}, nil
}