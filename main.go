package main

import (
	"fmt"
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

// Item is struct for DynamoDB
type Item struct {
	MyHashKey  string
	MyRangeKey int
	MyText     string
}

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

	/*
		CRUD DEMO
	*/

	// CREATE
	log.Println("START ~create~")
	item := Item{
		MyHashKey:  "MyPartitionKey",
		MyRangeKey: 1,
		MyText:     "My First Text",
	}
	if err := table.Put(item).Run(); err != nil {
		log.Printf("Failed to put item[%+v]\n", item)
	}
	log.Println("END ~create~")

	// READ
	log.Println("START ~read~")
	var readResult Item
	err = table.Get("MyHashKey", item.MyHashKey).Range("MyRangeKey", dynamo.Equal, item.MyRangeKey).One(&readResult)
	if err != nil {
		log.Println("Failed to get item")
	}
	fmt.Printf("*** Get item[%+v] ***\n", readResult)
	log.Println("END ~read~")

	// UPDATE
	log.Println("START ~update~")
	var updateResult Item
	err = table.Update("MyHashKey", item.MyHashKey).Range("MyRangeKey", item.MyRangeKey).Set("MyText", "My Second Text").Value(&updateResult)
	if err != nil {
		log.Printf("Failed to update item[%+v]\n", err)
	}
	log.Println("END ~update~")

	// DELETE
	log.Println("START ~delete~")
	err = table.Delete("MyHashKey", item.MyHashKey).Range("MyRangeKey", item.MyRangeKey).Run()
	if err != nil {
		log.Printf("Failed to delete item[%+v\n]", item)
	}
	log.Println("END ~delete~")

}
