package res

import (
	"log"
	service "sns-sqs/notification-service/services"
)

type ApiManager struct {
	Logger          *log.Logger
	SqsService      service.SqsHandler
	DynamoDbService service.DynamoDBHandler
}

func NewApiManager(logger *log.Logger, snsHandler service.SqsHandler, dynamoDbHandler service.DynamoDBHandler) *ApiManager {
	return &ApiManager{logger, snsHandler, dynamoDbHandler}
}
