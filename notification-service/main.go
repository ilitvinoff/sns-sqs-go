package main

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/joho/godotenv"
	"log"
	"sns-sqs/notification-service/res"
	service "sns-sqs/notification-service/services"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println(fmt.Sprintf("ERROR: can't load .env {%s}", err))
		return
	}

	cfg, err := LoadConfig()
	if err != nil {
		log.Println(fmt.Sprintf("ERROR: can't load cfg {%s}", cfg))
		return
	}

	sess, err := session.NewSessionWithOptions(session.Options{Config: aws.Config{Region: &cfg.AwsRegion}})
	if err != nil {
		log.Println(fmt.Sprintf("ERROR: can't create aws session {%s}", cfg))
		return
	}

	sqsSvc, err := service.NewSqsService(sess, cfg.QueueName, cfg.ReceiveMsgTimeout)
	if err != nil {
		log.Println(fmt.Sprintf("ERROR: can't create sqs service {%s}", err))
		return
	}

	ddbSvc := service.NewDynamoDbService(sess, cfg.DynamoDbName)

	svcManager := res.NewApiManager(log.Default(), sqsSvc, ddbSvc)

	for {
		svcManager.HandleMessage()
	}
}
