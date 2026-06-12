package main

import (
	"context"
	"fmt"
	"log"
	"time"

	article "github.com/maryam-nokohan/go-article/proto"
	"google.golang.org/grpc"
)

func main() {

	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("connection failed: %v", err)
	}
	defer conn.Close()

	client := article.NewArticleServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	stream, err := client.ProcessArticle(ctx)
	if err != nil {
		log.Fatalf("could not start stream: %v", err)
	}

	articles := []article.Article{
		{
			Title: "Go Microservices just another",
			Body:  "Go is a great language for building fast and scalable microservices.",
		},
		{
			Title: "Hexagonal Architecture design",
			Body:  "Hexagonal architecture isolates business logic from infrastructure.",
		},
	}

	for _, a := range articles {

		req := &article.ArticleRequest{
			Article: &article.Article{
				Title: a.Title,
				Body:  a.Body,
			},
		}

		fmt.Println("Sending article:", a.Title)

		err := stream.Send(req)
		if err != nil {
			log.Fatalf("send failed: %v", err)
		}
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("receive failed: %v", err)
	}

	fmt.Println("\nServer response:")
	fmt.Println("Created at:", res.CreatedAt)

}
