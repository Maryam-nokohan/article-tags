package main

import (
	"fmt"
	"log"

	"github.com/maryam-nokohan/go-article/internal/adapters/grpc"
	"github.com/maryam-nokohan/go-article/internal/adapters/mongo"
	"github.com/maryam-nokohan/go-article/internal/application"
)


func main(){
	fmt.Print("where")
	repo , err := mongo.NewMongoRepo()
	if err != nil {
		log.Fatal(err)
	}
	tagExtractor := application.NewTagEctractorService()
	articleService := application.NewArticleService(repo , tagExtractor)
	gprcAdaptor := grpc.NewServer(articleService)
	gprcAdaptor.Run(repo.Config.Port)

}

