package storage

import (
	envcfg "Douyin_Demo/config"
	"context"
	"encoding/json"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"io"
	"log"
	"net/http"
)

type LambdaResponseBody struct {
	ThumbnailFileName string `json:"thumbnailFileName"`
	Status            string `json:"status"`
}

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

// GetThumbnailLink asks lambda to create thumbnail and returns the link to the thumbnail of the object in the bucket.
func GetThumbnailLink(fileName string) (string, error) {
	lambdaFunctionUrl := envcfg.AppConfig.AWS.LambdaFunctionUrl

	response, err := http.Get(lambdaFunctionUrl + "?videoFileName=" + fileName + "&triggerBucketName=" + envcfg.AppConfig.AWS.BucketName)
	if err != nil {
		return "", err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return "", fmt.Errorf("failed to get thumbnail link from lambda function, status code: %v", response.StatusCode)
	}

	responseBody, err := io.ReadAll(response.Body)
	if err != nil {
		return "", err
	}

	var lambdaResponseBody LambdaResponseBody
	err = json.Unmarshal(responseBody, &lambdaResponseBody)
	if err != nil {
		return "", err
	}

	if lambdaResponseBody.Status != "success" {
		return "", fmt.Errorf("failed to get thumbnail link from lambda function, status: %v", lambdaResponseBody.Status)
	}

	thumbnailFileName := lambdaResponseBody.ThumbnailFileName

	return "https://" + envcfg.AppConfig.AWS.BucketName + ".s3." + envcfg.AppConfig.AWS.Region + ".amazonaws.com/" + thumbnailFileName, nil
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
