package main

// 参考了
// https://www.8kiz.cn/archives/24601.html
// https://github.com/minio/console/blob/63c6d8952bf148c20019c574da5dfa9b30c4d0cf/api/user_objects.go#L576
// 下面的代码是 stream 的，下载超过 1GiB 的 world 世界，占用内存不超过 20MiB

import (
	"archive/zip"
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	_ "net/http/pprof"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

const addr = ":5671"

const (
	endpoint        = "minio.570499536.xyz:3000"
	accessKeyID     = "uhziel@gmail.com"
	secretAccessKey = "o3XZ6RIGsXmCPa"
	useSSL          = true
	bucketName      = "sv7fj2nf2hmb7"
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
	http.HandleFunc("GET /download", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Disposition", "attachment; filename=hello.zip")
		w.Header().Add("Content-Type", "application/zip")

		zw := zip.NewWriter(w)
		defer zw.Close()

		chObjects := minioClient.ListObjects(ctx, bucketName, minio.ListObjectsOptions{
			Prefix:    prefix,
			Recursive: true,
		})

		for v := range chObjects {
			if v.Err != nil {
				log.Fatalln(v.Err)
				break
			}

			cw, err := zw.CreateHeader(&zip.FileHeader{
				Name:     v.Key,
				Method:   zip.Store,
				Modified: v.LastModified,
			})
			if err != nil {
				continue
			}

			obj, err := minioClient.GetObject(ctx, bucketName, v.Key, minio.GetObjectOptions{})
			if err != nil {
				fmt.Println("cannot open object:", v.Key)
				continue
			}

			_, err = io.Copy(cw, obj)
			obj.Close()
			if err != nil {
				fmt.Println("send object fail:", v.Key)
				continue
			}
		}
	})

	fmt.Println("listen at", addr)
	http.ListenAndServe(addr, nil)
}
