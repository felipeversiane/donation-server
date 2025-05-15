package filestorage

import (
	"fmt"
	"log/slog"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/felipeversiane/donation-server/config"
)

type fileStorage struct {
	client *s3.S3
	config config.FileStorageConfig
}

type Interface interface {
	Client() *s3.S3
	Bucket() string
	URL() string
}

func New(cfg config.FileStorageConfig) (Interface, error) {
	slog.Info("initializing file storage connection...")

	sess, err := session.NewSession(&aws.Config{
		Region:           aws.String(cfg.Region),
		Credentials:      credentials.NewStaticCredentials(cfg.AccessKey, cfg.SecretKey, ""),
		Endpoint:         aws.String(cfg.Endpoint),
		S3ForcePathStyle: aws.Bool(true),
	})
	if err != nil {
		slog.Info("unable to create S3 session", "error", err)
		return nil, fmt.Errorf("unable to create S3 session: %w", err)
	}

	client := s3.New(sess)

	fs := &fileStorage{
		client: client,
		config: cfg,
	}

	slog.Info("file storage initialized successfully")

	return fs, nil
}

func (f *fileStorage) Client() *s3.S3 {
	return f.client
}

func (f *fileStorage) Bucket() string {
	return f.config.Bucket
}

func (f *fileStorage) URL() string {
	return f.config.URL
}
