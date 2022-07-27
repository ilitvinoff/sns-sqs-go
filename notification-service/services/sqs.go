package service

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
)

type SqsHandler interface {
	ReceiveMessage() (*sqs.ReceiveMessageOutput, error)
	DeleteMessage(messageHandle *string) (*sqs.DeleteMessageOutput, error)
}

type SqsService struct {
	*sqs.SQS
	QueueName string
	QueueURL  string
	Timeout   int64
}

func NewSqsService(s *session.Session, name string, timeout int64) (*SqsService, error) {
	sqsSvc := sqs.New(s)
	urlResult, err := sqsSvc.GetQueueUrl(&sqs.GetQueueUrlInput{QueueName: &name})
	if err != nil {
		return nil, err
	}

	queueURL := urlResult.QueueUrl
	return &SqsService{sqsSvc, name, *queueURL, timeout}, nil
}

func (s *SqsService) ReceiveMessage() (*sqs.ReceiveMessageOutput, error) {
	return s.SQS.ReceiveMessage(&sqs.ReceiveMessageInput{
		AttributeNames: []*string{
			aws.String(sqs.MessageSystemAttributeNameSentTimestamp),
		},
		MessageAttributeNames: []*string{
			aws.String(sqs.QueueAttributeNameAll),
		},
		QueueUrl:            &s.QueueURL,
		MaxNumberOfMessages: aws.Int64(1),
		VisibilityTimeout:   &s.Timeout,
	})
}

func (s *SqsService) DeleteMessage(messageHandle *string) (*sqs.DeleteMessageOutput, error) {
	return s.SQS.DeleteMessage(&sqs.DeleteMessageInput{
		QueueUrl:      &s.QueueURL,
		ReceiptHandle: messageHandle,
	})
}
