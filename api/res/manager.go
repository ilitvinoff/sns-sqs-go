package res

import (
	"log"
	service "sns-sqs/api/services"
)

type ApiManager struct {
	Logger     *log.Logger
	SnsService service.SnsHandler
}

func NewApiService(logger *log.Logger, snsHandler service.SnsHandler) *ApiManager {
	return &ApiManager{logger, snsHandler}
}
