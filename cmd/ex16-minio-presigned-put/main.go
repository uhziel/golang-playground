package main

import (
	"context"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"log"
	"time"
)

const (
	endpoint        = "play.min.io"
	accessKeyID     = "Q3AM3UQ867SPQQA43P2F"
	secretAccessKey = "zuf+tfteSlswRu7BJ86wekitnifILbZam1KYY3TG"
	useSSL          = true
	bucketName      = "uhziel"
	location        = "us-east-1"
	objectName      = "main.go"
	filePath        = "./main.go"
	contentType     = "text/plain"
)

func main() {
	client, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
		Secure: useSSL,
	})

	if err != nil {
		log.Fatalln(err)
	}

	ctx := context.Background()
	err = client.MakeBucket(ctx, bucketName, minio.MakeBucketOptions{
		Region: location,
	})
	if err != nil {
		exists, errBucketExists := client.BucketExists(ctx, bucketName)
		if exists && errBucketExists == nil {
			log.Println("We already own", bucketName)
		} else {
			log.Fatalln(err)
		}
	} else {
		log.Println("Successfully created", bucketName)
	}

	presignedUrl, err := client.PresignedPutObject(ctx, bucketName, objectName, time.Duration(300)*time.Second)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("upload cmd:")
	log.Printf("curl -X PUT -H 'Content-Type: text/plain' --data-binary '@main.go' '%s'", presignedUrl)
}
