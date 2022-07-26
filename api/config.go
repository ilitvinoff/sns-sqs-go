package main

import (
	"fmt"
	"os"
)

type Config struct {
	SnsArn    string
	AwsRegion string
}

func LoadConfig() (*Config, error) {
	arn, ok := os.LookupEnv("SNS_ARN")
	if !ok {
		return nil, fmt.Errorf("check if SNS_ARN defined")
	}

	region, ok := os.LookupEnv("AWS_REGION")
	if !ok {
		return nil, fmt.Errorf("check if AWS_REGION defined")
	}

	return &Config{arn, region}, nil
}