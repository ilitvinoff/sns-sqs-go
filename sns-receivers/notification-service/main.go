package main

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/joho/godotenv"
	"log"
	service "sns-sqs/common/services"
	"sns-sqs/sns-receivers/notification-service/res"
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

	svcManager := res.NewManager(log.Default(), sqsSvc, ddbSvc)

	for {
		svcManager.HandleMessage()
	}
}
