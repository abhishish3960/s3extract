package services

import (
	"io"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

type S3Service struct {
	client *s3.S3
	bucket string
}

func NewS3Service(bucket string) *S3Service {
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("ap-south-1"),
	})
	if err != nil {
		panic("Failed to create AWS session")
	}
	return &S3Service{
		client: s3.New(sess),
		bucket: bucket,
	}
}

// GeneratePresignedURL generates a presigned URL for the specified key
func (s *S3Service) GeneratePresignedURL(key string) (string, error) {
	req, _ := s.client.GetObjectRequest(&s3.GetObjectInput{
		Bucket: aws.String(s.bucket),
		Key:    aws.String(key),
	})
	return req.Presign(15 * time.Minute)
}

// ListFiles lists the files in the specified prefix
func (s *S3Service) ListFiles(prefix string) ([]string, error) {
	var keys []string
	err := s.client.ListObjectsV2Pages(&s3.ListObjectsV2Input{
		Bucket: aws.String(s.bucket),
		Prefix: aws.String(prefix),
	}, func(page *s3.ListObjectsV2Output, lastPage bool) bool {
		for _, obj := range page.Contents {
			keys = append(keys, *obj.Key)
		}
		return !lastPage
	})
	if err != nil {
		return nil, err
	}
	return keys, nil
}

// GetFileContent retrieves the content of a text file from S3
func (s *S3Service) GetFileContent(key string) (string, error) {
	resp, err := s.client.GetObject(&s3.GetObjectInput{
		Bucket: aws.String(s.bucket),
		Key:    aws.String(key),
	})
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(body), nil
}
