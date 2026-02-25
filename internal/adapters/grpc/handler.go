package grpc

import (
	"context"
	"io"
	"log"
	"net"
	"runtime"
	"time"

	"github.com/maryam-nokohan/go-article/internal/application"
	"github.com/maryam-nokohan/go-article/internal/domain"
	"github.com/maryam-nokohan/go-article/internal/pkg"
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

func (s *Server) ProcessArticle(stream article.ArticleService_ProcessArticleServer) error {
	log.Println("handler : Received a new article stream")

	pool := pkg.New(runtime.NumCPU())
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()


	pool.Run(ctx, func(a *domain.Article) error {
		if err := s.service.ProcessArticle(a); err != nil {
			log.Println("Error saving article:", err)
			return err
		}
		return nil
	})

	for {
		req, err := stream.Recv()
		if err == io.EOF {
			pool.Close()

			return stream.SendAndClose(&article.ArticleResponse{
				CreatedAt: time.Now().Format(time.RFC3339),
			})
		}
		if err != nil {
			pool.Close()
			return status.Errorf(codes.Internal, "recv error: %v", err)
		}
		if req.Article == nil {
			continue
		}
		a := &domain.Article{
			Title: req.Article.Title,
			Body:  req.Article.Body,
		}
		pool.Submit(a)
	}
}

func (s *Server) TopTags(ctx context.Context, req *article.TopTagsRequst) (*article.TopTagResponse, error) {
	TopN := req.GetTopn()
	log.Println("handler : Received a request for top tags with limit:", TopN)
	tags, err := s.service.GetTopTags(ctx, TopN)
	if err != nil {
		return nil, err
	}

	grpcTags := make([]*article.Tag,0, len(tags))
	for _, t := range tags {
		if t.Word == " " || t.Word == ""  || t.Freq == 0{
			continue
		}
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
	lis, err := net.Listen("tcp", ":"+address)
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
