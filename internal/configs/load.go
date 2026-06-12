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
	nats_url := os.Getenv("NATS_URL")
	endpoint   := os.Getenv("ENDPOINT")  
	bucket := os.Getenv("BUCKET")
	region     := os.Getenv("REGION")    
	accesskey  := os.Getenv("ACCESSKEY") 
	secretkey := os.Getenv("SECRETEKEY")

	if uri == "" || nats_url== ""|| bucket== ""|| dbName == "" || gRPCPort == "" || endpoint == "" || region == "" || accesskey == "" || secretkey==""{
		log.Fatal("setup your .env")
	}

	return &Config{
		DBName: dbName,
		GRPC_Port: gRPCPort,
		URI:    uri,
		NATS_URL: nats_url,
		BUCKET : bucket,
		ENDPOINT: endpoint,
		REGION: region,
		ACCESSKEY: accesskey,
		SECRETEKEY: secretkey,
	}, nil
}