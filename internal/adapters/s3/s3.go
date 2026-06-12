package s3

import (
	"bytes"
	"context"
	"fmt"
	"io"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

type S3Storage struct {
	client *s3.S3
	bucket string
}

func NewS3Storage(endpoint, region, bucket, accesskey, secretkey string) (*S3Storage, error) {
	sess, err := session.NewSession(&aws.Config{
		Endpoint: aws.String(endpoint),
		Region:   aws.String(region),
		Credentials: credentials.NewStaticCredentials(
			accesskey,
			secretkey,
			"",
		),
	})
	if err != nil {
		return nil, fmt.Errorf("Failed to create session %v", err)
	}
	client := s3.New(sess)
	_, err = client.CreateBucket(&s3.CreateBucketInput{
		Bucket: aws.String(bucket),
	})
	if err != nil {
		fmt.Printf("Note: Bucket may already exist: %v\n", err)
	}
	return &S3Storage{
		client: client,
		bucket: bucket,
	}, nil
}
func (s *S3Storage) Upload(key string, data []byte) (string, error) {
	_, err := s.client.PutObjectWithContext(context.TODO(), &s3.PutObjectInput{
		Bucket:      &s.bucket,
		Key:         aws.String(key),
		Body:        bytes.NewReader(data),
		ContentType: aws.String("application/json"),
	})
	if err != nil {
		return "", fmt.Errorf("failed to upload: %w", err)
	}

	return key, nil
}

func (s *S3Storage) Download(key string) ([]byte, error) {
	result, err := s.client.GetObjectWithContext(context.Background(), &s3.GetObjectInput{
		Bucket: aws.String(s.bucket),
		Key:    aws.String(key),
	})
	if err != nil {
		return nil, fmt.Errorf("failed to download: %w", err)
	}
	defer result.Body.Close()

	data, err := io.ReadAll(result.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read data: %w", err)
	}

	return data, nil
}
