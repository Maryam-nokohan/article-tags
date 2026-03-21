package main

import (
	"log"

	"github.com/maryam-nokohan/go-article/internal/adapters/grpc"
	"github.com/maryam-nokohan/go-article/internal/adapters/mongo"
	"github.com/maryam-nokohan/go-article/internal/adapters/nats"
	"github.com/maryam-nokohan/go-article/internal/adapters/s3"
	application "github.com/maryam-nokohan/go-article/internal/application"
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

	broker, err := nats.NewNatsBroker(config.NATS_URL)
	if err != nil {
		log.Fatal(err)
	}

	storage, err := s3.NewS3Storage(
		config.ENDPOINT,
		config.REGION,
		config.BUCKET,
		config.ACCESSKEY,
		config.SECRETEKEY,
	)
	if err != nil {
		log.Fatal(err)
	}

	tagExtractor := application.NewTagExtractorService()

	ingestionService := application.NewIngestionService(broker, storage)

	processingService := application.NewProcessingService(
		storage,
		repo,
		tagExtractor,
	)

	err = broker.Subscribe("article.created", func(data []byte) error {
		return processingService.HandleArticleCreated(data)
	})
	if err != nil {
		log.Fatal(err)
	}

	grpcAdaptor := grpc.NewServer(
		ingestionService,
		processingService,
	)

	err = grpcAdaptor.Run(config.GRPC_Port)
	if err != nil {
		log.Fatal(err)
	}
}
