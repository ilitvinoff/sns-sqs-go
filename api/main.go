package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"os"
	"sns-sqs/api/res"
	svc "sns-sqs/api/services"
)

func main() {
	arn, ok := os.LookupEnv("SNS_ARN")
	if !ok {
		log.Println("ERROR: check if SNS_ARN defined")
	}

	snsSvc, err := svc.NewSnsService(arn)
	if err != nil {
		log.Println(fmt.Sprintf("ERROR: can't init sns service {%s}", err))
	}

	svcManager := res.NewApiService(log.Default(), snsSvc)

	engine := gin.New()
	engine.POST("/data-to-manage", svcManager.Publish)
}
