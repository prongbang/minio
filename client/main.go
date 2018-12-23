package main

import (
	"log"
	"net/url"
	"time"

	minio "github.com/minio/minio-go"
)

func main() {
	endpoint := "localhost:9001"
	accessKeyID := "6DVY3Pkc4z"
	secretAccessKey := "FAAmZ0Evr7"
	useSSL := false

	// Initialize minio client object.
	minioClient, err := minio.New(endpoint, accessKeyID, secretAccessKey, useSSL)
	if err != nil {
		log.Fatalln(err)
	}

	// Make a new bucket called anim.
	bucketName := "anim"
	location := "us-east-1"

	err = minioClient.MakeBucket(bucketName, location)
	if err != nil {
		// Check to see if we already own this bucket (which happens if you run this twice)
		exists, err := minioClient.BucketExists(bucketName)
		if err == nil && exists {
			log.Printf("We already own %s\n", bucketName)
		} else {
			log.Fatalln(err)
		}
	} else {
		log.Printf("Successfully created %s\n", bucketName)
	}

	// Upload the image file
	objectName := "minio.png"
	filePath := "./image/minio.png"
	contentType := "application/octet-stream"

	// Upload the image file with FPutObject
	n, err := minioClient.FPutObject(bucketName, objectName, filePath, minio.PutObjectOptions{ContentType: contentType})
	if err != nil {
		log.Fatalln(err)
	}
	log.Printf("Successfully uploaded %s of size %d\n", objectName, n)

	// Set request parameters
	reqParams := make(url.Values)

	// Gernerate presigned get object url.
	presignedURL, err := minioClient.PresignedGetObject(bucketName, objectName, time.Duration(1000)*time.Second, reqParams)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(presignedURL)
}
