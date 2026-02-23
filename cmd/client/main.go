package main

import (
	"context"
	"fmt"
	"log"
	"time"

	article "github.com/maryam-nokohan/go-article/proto"
	"google.golang.org/grpc"
)

func DoClientStreaming(c article.ArticleServiceClient) {
	fmt.Println("Client streaming ...(writing articles to DB)...")
	stream, err := c.ProcessArticle(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	requests := []*article.ArticleRequest{
			{
			Article: &article.Article{
				Title: "The Rise of Artificial Intelligence",
				Body: `Artificial intelligence has transformed modern technology. 
				Machine learning and deep learning systems are now used in healthcare, 
				finance, cybersecurity, and autonomous vehicles. 
				Companies invest heavily in AI research to build scalable intelligent systems.`,
			},
		},
		{
			Article: &article.Article{
				Title: "Cloud Computing and Distributed Systems",
				Body: `Cloud computing enables scalable distributed systems. 
				Microservices architecture improves maintainability and performance. 
				Technologies like Docker and Kubernetes simplify deployment and orchestration.`,
			},
		},
		{
			Article: &article.Article{
				Title: "The Future of Cybersecurity",
				Body: `Cybersecurity threats are increasing globally. 
				Encryption, zero-trust architecture, and secure authentication 
				are critical components of modern software systems.`,
			},
		},
	}

	for _, req := range requests {
		log.Printf("Sending request: %v", req)
		if err := stream.Send(req); err != nil {
			log.Fatal(err)
		}
	}

	resp, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Server response tags:")
	for _, t := range resp.GetTags() {
		fmt.Printf(" - %s : %d\n", t.Word, t.Freq)
	}
}

func DoUnary(c article.ArticleServiceClient) {
	fmt.Println("Unary call ...(top tags)")

	req := &article.TopTagsRequst{
		Topn: 3,
	}

	res, err := c.TopTags(context.Background(), req)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Top tags:")
	for _, t := range res.GetTags() {
		fmt.Printf(" - %s : %d\n", t.Word, t.Freq)
	}
}

func main() {
	// Dial the gRPC server
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure(), grpc.WithBlock(), grpc.WithTimeout(5*time.Second))
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	client := article.NewArticleServiceClient(conn)

	DoClientStreaming(client)
	DoUnary(client)
}