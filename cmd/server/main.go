package main

import (
	"log"

	"github.com/maryam-nokohan/go-article/internal/adapters/grpc"
	"github.com/maryam-nokohan/go-article/internal/adapters/mongo"
	"github.com/maryam-nokohan/go-article/internal/application"
	"github.com/maryam-nokohan/go-article/internal/configs"
)

func main() {
	config, err := configs.Newconfig()
	if err != nil {
		log.Fatal(err)
	}

	repo, err := mongo.NewMongoRepo(config)
	if err != nil {
		log.Fatal(err)
	}
	tagExtractor := application.NewTagExtractorService()
	articleService := application.NewArticleService(repo, tagExtractor)
	grpcAdaptor := grpc.NewServer(articleService)
	err = grpcAdaptor.Run(config.GRPC_Port)
	if err != nil {
		log.Fatal(err)
	}
}
