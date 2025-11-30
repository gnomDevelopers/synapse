package minios3

import (
	"context"
	"fmt"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"log"
	"net/url"
	"synapse/internal/config"
	"time"
)

type S3 struct {
	S3 *minio.Client
}

func MinioConnection(conf *config.Config) *S3 {
	endpoint := conf.Minio.Endpoint               // адрес с3
	accessKeyID := conf.Minio.AccessKeyID         // логин
	secretAccessKey := conf.Minio.SecretAccessKey // пароль
	useSSL := conf.Minio.UseSSL

	minioClient, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
		Secure: useSSL,
		Region: "ru-central-1",
	})
	if err != nil {
		log.Fatalln(fmt.Errorf("minio client new error: %w", err))
	}

	location := "ru-central-1"

	err = minioClient.MakeBucket(context.Background(), "study-material", minio.MakeBucketOptions{Region: location})
	if err != nil {

		exists, errBucketExists := minioClient.BucketExists(context.Background(), "study-material")
		if errBucketExists == nil && exists {
			fmt.Printf("We already own %s\n", "study-material")
		} else {
			log.Fatalln(fmt.Errorf("minio MakeBucket study-material error: %w", err))
		}
	} else {
		fmt.Printf("successfully created %s\n", "study-material")
	}

	err = minioClient.MakeBucket(context.Background(), "assignment-media", minio.MakeBucketOptions{Region: location})
	if err != nil {

		exists, errBucketExists := minioClient.BucketExists(context.Background(), "task-media")
		if errBucketExists == nil && exists {
			fmt.Printf("We already own %s\n", "task-media")
		} else {
			log.Fatalln(fmt.Errorf("minio MakeBucket task-media error: %w", err))
		}
	} else {
		fmt.Printf("successfully created %s\n", "assignment-media")
	}

	return &S3{minioClient}
}

func (s3 *S3) FPutObject(ctx context.Context, bucketName, fileName, filePath, contentType string) error {
	_, err := s3.S3.FPutObject(ctx, bucketName, fileName, filePath, minio.PutObjectOptions{ContentType: contentType})
	if err != nil {
		return err
	}
	return nil
}

func (s3 *S3) PresignedGetObject(ctx context.Context, bucketName, filePath string, expiration time.Duration) (*url.URL, error) {
	object, err := s3.S3.PresignedGetObject(ctx, bucketName, filePath, expiration, nil)
	if err != nil {
		return nil, err
	}
	return object, nil
}

func (s3 *S3) GetObject(ctx context.Context, bucketName, filePath string) (*minio.Object, error) {
	object, err := s3.S3.GetObject(ctx, bucketName, filePath, minio.GetObjectOptions{})
	if err != nil {
		return nil, err
	}
	return object, nil
}

func (s3 *S3) RemoveObject(ctx context.Context, bucketName, fileName string) error {
	err := s3.S3.RemoveObject(ctx, bucketName, fileName, minio.RemoveObjectOptions{})
	if err != nil {
		return err
	}
	return nil
}
