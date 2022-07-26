package service

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sns"
)

type SnsHandler interface {
	Publish(msg string) (*sns.PublishOutput, error)
}

type SnsService struct {
	*sns.SNS
	Arn string
}

func NewSnsService(arn string) (*SnsService, error) {
	currentSession, err := session.NewSessionWithOptions(session.Options{SharedConfigState: session.SharedConfigEnable})
	if err != nil {
		return nil, err
	}
	return &SnsService{sns.New(currentSession), arn}, nil
}

func (s *SnsService) Publish(msg string) (*sns.PublishOutput, error) {
	return s.SNS.Publish(&sns.PublishInput{
		Message:  &msg,
		TopicArn: &s.Arn,
	})
}
