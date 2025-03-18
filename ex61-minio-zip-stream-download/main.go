package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

const addr = ":5671"

const (
	endpoint        = "minio.570499536.xyz:3000"
	accessKeyID     = "uhziel@gmail.com"
	secretAccessKey = "o3XZ6RIGsXmCPa"
	useSSL          = true
	bucketName      = "test-server-agent"
	prefix          = "world"
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
}

func main() {
	http.HandleFunc("POST /download", func(w http.ResponseWriter, r *http.Request) {
		chObjects := minioClient.ListObjects(ctx, bucketName, minio.ListObjectsOptions{
			Prefix:    prefix,
			Recursive: true,
		})

		for v := range chObjects {
			fmt.Println(v.Key)
		}
	})

	fmt.Println("listen at", addr)
	http.ListenAndServe(addr, nil)
}
