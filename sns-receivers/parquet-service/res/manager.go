package res

import (
	"log"
	service "sns-sqs/common/services"
)

type ApiManager struct {
	Logger           *log.Logger
	SqsService       service.SqsHandler
	DynamoDbService  service.DynamoDBHandler
	S3ParquetService service.S3ParquetHandler
}

func NewManager(
	logger *log.Logger,
	snsHandler service.SqsHandler,
	dynamoDbHandler service.DynamoDBHandler,
	s3parquetService service.S3ParquetHandler,
) *ApiManager {
	return &ApiManager{
		logger,
		snsHandler,
		dynamoDbHandler,
		s3parquetService,
	}
}
