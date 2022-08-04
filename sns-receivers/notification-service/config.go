package main

import (
	"fmt"
	"os"
	"strconv"
)

type Config struct {
	QueueName         string
	DynamoDbName      string
	AwsRegion         string
	ReceiveMsgTimeout int64
}

func LoadConfig() (*Config, error) {
	name, ok := os.LookupEnv("QUEUE_NAME")
	if !ok {
		return nil, fmt.Errorf("check if QUEUE_NAME defined")
	}

	dynamodbTableName, ok := os.LookupEnv("DYNAMODB_TABLE_NAME")
	if !ok {
		return nil, fmt.Errorf("check if DYNAMODB_TABLE_NAME defined")
	}

	timeout, ok := os.LookupEnv("RECEIVE_MSG_TIMEOUT")
	if !ok {
		return nil, fmt.Errorf("check if RECEIVE_MSG_TIMEOUT defined")
	}

	timeoutInt, err := strconv.Atoi(timeout)
	if err != nil {
		return nil, fmt.Errorf("can't parse timeout: {%s}", err)
	}

	region, ok := os.LookupEnv("AWS_REGION")
	if !ok {
		return nil, fmt.Errorf("check if AWS_REGION defined")
	}

	return &Config{
		QueueName:         name,
		DynamoDbName:      dynamodbTableName,
		AwsRegion:         region,
		ReceiveMsgTimeout: int64(timeoutInt),
	}, nil
}
