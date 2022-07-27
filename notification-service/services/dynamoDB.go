package service

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"sns-sqs/notification-service/models"
)

type DynamoDBHandler interface {
	IsItemExists(id string) (bool, error)
	PutMessage(id string) error
}

type DynamoDbService struct {
	*dynamodb.DynamoDB
	TableName string
}

func NewDynamoDbService(sess *session.Session, tableName string) *DynamoDbService {
	return &DynamoDbService{dynamodb.New(sess), tableName}
}

func (s *DynamoDbService) IsItemExists(id string) (bool, error) {
	result, err := s.GetItem(&dynamodb.GetItemInput{
		TableName: aws.String(s.TableName),
		Key: map[string]*dynamodb.AttributeValue{
			"id": {
				S: aws.String(id),
			},
		},
	})
	if err != nil {
		return false, err
	}

	return result.Item != nil, nil
}

func (s *DynamoDbService) PutMessage(id string) error {
	item := &models.Item{ID: id}
	av, err := dynamodbattribute.MarshalMap(item)
	if err != nil {
		return err
	}

	input := &dynamodb.PutItemInput{
		Item:      av,
		TableName: aws.String(s.TableName),
	}

	_, err = s.PutItem(input)
	if err != nil {
		return err
	}

	return nil
}
