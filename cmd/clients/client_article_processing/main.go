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

	res, err := client.TopTags(ctx, &article.TopTagsRequst{Topn: 5})
	if err != nil {
		log.Fatal(err)
	}

	for _, t := range res.Tags {
		fmt.Println(t.Word, t.Freq)
	}

}
