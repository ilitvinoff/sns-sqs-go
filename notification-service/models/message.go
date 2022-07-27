package models

import "github.com/aws/aws-sdk-go/service/sqs"

type AwsMessages []*sqs.Message

type Message struct {
	ID  string `json:"id" binding:"required"`
	Msg string `json:"msg" binding:"required"`
}
