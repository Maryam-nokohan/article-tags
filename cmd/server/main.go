package main

import (
	"log"

	"github.com/maryam-nokohan/go-article/internal/adapters/grpc"
	"github.com/maryam-nokohan/go-article/internal/adapters/mongo"
	"github.com/maryam-nokohan/go-article/internal/application"
)

func main(){
	repo , err := mongo.NewMongoRepo()
	if err != nil {
		log.Fatal(err)
	}
	tagExtractor := application.NewTagExtractorService()
	articleService := application.NewArticleService(repo , tagExtractor)
	gprcAdaptor := grpc.NewServer(articleService)
	err = gprcAdaptor.Run(repo.Config.GRPC_Port)
	if err != nil {
		log.Fatal(err)
	}
}

