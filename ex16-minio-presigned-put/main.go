package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

const (
	endpoint        = "minio.570499536.xyz:3000"
	accessKeyID     = "uhziel@gmail.com"
	secretAccessKey = "o3XZ6RIGsXmCPa"
	useSSL          = true
	bucketName      = "uhziel"
	location        = "us-east-1"
	objectName      = "main.go"
)

var ctx = context.Background()
var minioClient *minio.Client

func init() {
	var err error
	minioClient, err = minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
		Secure: useSSL,
	})

	if err != nil {
		log.Fatalln(err)
	}

	err = minioClient.MakeBucket(ctx, bucketName, minio.MakeBucketOptions{
		Region: location,
	})
	if err != nil {
		exists, errBucketExists := minioClient.BucketExists(ctx, bucketName)
		if exists && errBucketExists == nil {
			log.Println("We already own", bucketName)
		} else {
			log.Fatalln(err)
		}
	} else {
		log.Println("Successfully created", bucketName)
	}
}

func s3PresignHandler(w http.ResponseWriter, r *http.Request) {
	presignedUrl, err := minioClient.PresignedPutObject(ctx, bucketName, objectName, time.Duration(300)*time.Second)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	io.WriteString(w, fmt.Sprintf("curl -X PUT -H 'Content-Type: text/plain' --data-binary '@%s' '%s'", objectName, presignedUrl))
}

func main() {
	http.HandleFunc("GET /presign", s3PresignHandler)
	http.HandleFunc("GET /{$}", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "index.html")
	})

	const addr = ":4567"
	log.Println("listen at", addr)
	log.Fatal(http.ListenAndServe(addr, nil))
}
