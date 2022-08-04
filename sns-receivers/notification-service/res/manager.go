package res

import (
	"log"
	service "sns-sqs/common/services"
)

type ApiManager struct {
	Logger          *log.Logger
	SqsService      service.SqsHandler
	DynamoDbService service.DynamoDBHandler
}

func NewManager(logger *log.Logger, sqsHandler service.SqsHandler, dynamoDbHandler service.DynamoDBHandler) *ApiManager {
	return &ApiManager{logger, sqsHandler, dynamoDbHandler}
}
