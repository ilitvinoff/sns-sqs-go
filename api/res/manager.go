package res

import (
	"log"
	service "sns-sqs/common/services"
)

type ApiManager struct {
	Logger     *log.Logger
	SnsService service.SnsHandler
}

func NewApiManager(logger *log.Logger, snsHandler service.SnsHandler) *ApiManager {
	return &ApiManager{logger, snsHandler}
}
