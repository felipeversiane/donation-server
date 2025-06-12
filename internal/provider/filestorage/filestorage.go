package filestorage

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/felipeversiane/donation-server/config"
	"github.com/felipeversiane/donation-server/pkg/logger"
)

type fileStorage struct {
	client *s3.S3
	config config.FileStorage
	logger logger.Interface
}

type Interface interface {
	Client() *s3.S3
	Bucket() string
	URL() string
	CreateBucket() error
}

func New(cfg config.FileStorage, logger logger.Interface) (Interface, error) {
	logger.Logger().Info("initializing file storage connection...")

	sess, err := session.NewSession(&aws.Config{
		Region:           aws.String(cfg.Region),
		Credentials:      credentials.NewStaticCredentials(cfg.AccessKey, cfg.SecretKey, ""),
		Endpoint:         aws.String(cfg.Endpoint),
		S3ForcePathStyle: aws.Bool(true),
	})
	if err != nil {
		logger.Logger().Error("unable to create S3 session", "error", err)
		return nil, fmt.Errorf("unable to create S3 session: %w", err)
	}

	client := s3.New(sess)

	fs := &fileStorage{
		client: client,
		config: cfg,
		logger: logger,
	}

	logger.Logger().Info("file storage initialized successfully")

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

func (f *fileStorage) CreateBucket() error {
	client := f.Client()
	bucket := f.Bucket()

	_, err := client.HeadBucket(&s3.HeadBucketInput{
		Bucket: aws.String(bucket),
	})

	if err == nil {
		f.logger.Logger().Info("bucket already exists", "bucket", bucket)
		return nil
	}

	if aerr, ok := err.(awserr.Error); ok && aerr.Code() == "NotFound" {
		f.logger.Logger().Warn("bucket not found, creating", "bucket", bucket)
		_, err = client.CreateBucket(&s3.CreateBucketInput{
			Bucket: aws.String(bucket),
		})
		if err != nil {
			f.logger.Logger().Error("failed to create bucket", "bucket", bucket, "error", err)
			return fmt.Errorf("failed to create bucket %s: %w", bucket, err)
		}
		f.logger.Logger().Info("bucket created successfully", "bucket", bucket)
		return nil
	}

	return fmt.Errorf("error checking bucket existence: %w", err)
}
