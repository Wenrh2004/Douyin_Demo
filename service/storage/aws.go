package storage

import (
	envcfg "Douyin_Demo/config"
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"io"
	"log"
)

var s3Client *s3.Client

func init() {
	sdkConfig, err := config.LoadDefaultConfig(context.TODO(),
		config.WithRegion(envcfg.AppConfig.AWS.Region),
		config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(envcfg.AppConfig.AWS.AccessKey, envcfg.AppConfig.AWS.Secret, "")),
	)
	if err != nil {
		fmt.Println("Couldn't load default configuration. Have you set up your AWS account?")
		log.Fatal(err)
		return
	}
	fmt.Println(envcfg.AppConfig.AWS.AccessKey, envcfg.AppConfig.AWS.Secret)

	s3Client = s3.NewFromConfig(sdkConfig)
}

// UploadFile reads from a file and puts the data into an object in a bucket.
func UploadFile(content io.Reader, fileName string) (*s3.PutObjectOutput, error) {
	output, err := s3Client.PutObject(context.TODO(), &s3.PutObjectInput{
		Bucket: &envcfg.AppConfig.AWS.BucketName,
		Key:    &fileName,
		Body:   content,
	})
	return output, err
}

// GetObjectLink returns the link to the object in the bucket.
func GetObjectLink(fileName string) string {
	return "https://" + envcfg.AppConfig.AWS.BucketName + ".s3." + envcfg.AppConfig.AWS.Region + ".amazonaws.com/" + fileName
}

func main() {

	sdkConfig, err := config.LoadDefaultConfig(context.TODO(),
		config.WithRegion(envcfg.AppConfig.AWS.Region),
		config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(envcfg.AppConfig.AWS.AccessKey, envcfg.AppConfig.AWS.Secret, "")),
	)
	if err != nil {
		fmt.Println("Couldn't load default configuration. Have you set up your AWS account?")
		log.Fatal(err)
		return
	}
	fmt.Println(envcfg.AppConfig.AWS.AccessKey, envcfg.AppConfig.AWS.Secret)

	s3Client = s3.NewFromConfig(sdkConfig)
	count := 10
	fmt.Printf("Let's list up to %v buckets for your account.\n", count)
	result, err := s3Client.ListBuckets(context.TODO(), &s3.ListBucketsInput{})
	if err != nil {
		fmt.Printf("Couldn't list buckets for your account. Here's why: %v\n", err)
		return
	}
	if len(result.Buckets) == 0 {
		fmt.Println("You don't have any buckets!")
	} else {
		if count > len(result.Buckets) {
			count = len(result.Buckets)
		}
		for _, bucket := range result.Buckets[:count] {
			fmt.Printf("\t%v\n", *bucket.Name)
		}
	}
}
