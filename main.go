package main

import (
	"log"
	"os"

	"github.com/guregu/dynamo"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/joho/godotenv"
)

const (
	TokyoResion = "ap-northeast-1"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// config client
	dynamoDbRegion := os.Getenv("AWS_REGION")
	disableSsl := false

	// ※dynamodb-localを利用する場合はEndpintのURLを設定する
	dynamoDbEndpoint := os.Getenv("DYNAMO_ENDPOINT")
	if len(dynamoDbEndpoint) != 0 {
		disableSsl = true
	}

	// defaltで東京リージョンを指定
	if len(dynamoDbRegion) == 0 {
		dynamoDbRegion = TokyoResion
	}

	db := dynamo.New(session.New(), &aws.Config{
		Region:     aws.String(dynamoDbRegion),
		Endpoint:   aws.String(dynamoDbEndpoint),
		DisableSSL: aws.Bool(disableSsl),
	})

	table := db.Table("MyFirstTable")
	_ = table
}
