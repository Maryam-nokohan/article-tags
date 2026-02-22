package grpc

import (
	"context"
	"io"
	"log"
	"net"
	"time"

	"github.com/maryam-nokohan/go-article/internal/application"
	"github.com/maryam-nokohan/go-article/internal/domain"
	article "github.com/maryam-nokohan/go-article/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Server struct {
	article.UnimplementedArticleServiceServer
	service    *application.ArticleService
	grpcServer *grpc.Server
}

func NewServer(articleService *application.ArticleService) *Server {
	s := &Server{
		service:    articleService,
		grpcServer: grpc.NewServer(),
	}
	article.RegisterArticleServiceServer(s.grpcServer, s)
	return s
}

func (s *Server) ProcessArticle(stream article.ArticleService_ProcessArticlesServer) error {
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&article.ArticleResponse{
				Article:   &article.Article{},
				Tags:      nil,
				CreatedAt: time.Now().Format(time.RFC3339),
			})
		}
		if err != nil {
			return status.Errorf(codes.Internal, "recv error: %v", err)
		}

		domainArticle := &domain.Article{
			Title: req.Article.Title,
			Body:  req.Article.Body,
		}

		go func(a *domain.Article) {
			if err := s.service.ProcessArticle(*a); err != nil {
				log.Println("Error saving article:", err)
			}
		}(domainArticle)
	}
}

func (s *Server) GetTopN(ctx context.Context, req *article.TopTagsRequst) (*article.TopTagResponse, error) {
	TopN := req.GetTopn()
	tags, err := s.service.GetTopTags(ctx, TopN)
	if err != nil {
		return nil, err
	}

	grpcTags := make([]*article.Tag, len(tags))
	for _, t := range tags {
		grpcTags = append(grpcTags, &article.Tag{
			Word: t.Word,
			Freq: t.Freq,
		})
	}
	return &article.TopTagResponse{
		Tags: grpcTags,
	}, nil
}
func (s *Server) Run(address string) error {
	log.Printf("Running the gRPC server on port %s...", address)
	lis, err := net.Listen("tcp", address)
	if err != nil {
		return err
	}
	if err := s.grpcServer.Serve(lis); err != nil {
		return err
	}
	return nil
}
func (s *Server) GracefulShutdown() {
	s.grpcServer.GracefulStop()
}
