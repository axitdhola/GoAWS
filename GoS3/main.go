// package main

// import (
// 	"context"
// 	"fmt"

// 	"github.com/aws/aws-sdk-go-v2/config"
// 	"github.com/aws/aws-sdk-go-v2/service/s3"
// )

// func main() {
// 	sdkConfig, err := config.LoadDefaultConfig(context.TODO())
// 	if err != nil {
// 		fmt.Println("Couldn't load default configuration")
// 		fmt.Println(err)
// 		return
// 	}

// 	s3Client := s3.NewFromConfig(sdkConfig)
// 	result, err := s3Client.ListBuckets(context.TODO(), &s3.ListBucketsInput{})
// 	if err != nil {
// 		fmt.Printf("Couldn't list buckets for your account. Error: %v\n", err)
// 		return
// 	}

// 	count := len(result.Buckets)
// 	for _, bucket := range result.Buckets[:count] {
// 		fmt.Printf("\t%v\n", *bucket.Name)
// 	}
// }
