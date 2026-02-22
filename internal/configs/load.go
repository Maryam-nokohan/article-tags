package configs

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func Newconfig() (*Config , error){
	if err := godotenv.Load();err != nil {
		return nil , err
	}

	dbName := os.Getenv("DB_NAME")
	port := os.Getenv("PORT")
	uri := os.Getenv("URI")

	if uri == "" || dbName == "" || port == "" {
		log.Fatal("setup your .env")
	}

	return &Config{
		DBName: dbName,
		Port: port,
		URI: uri,
	} , nil
}