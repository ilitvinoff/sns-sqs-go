package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
	"sns-sqs/api/res"
	svc "sns-sqs/common/services"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println(fmt.Sprintf("ERROR: can't load .env {%s}", err))
		return
	}

	cfg, err := LoadConfig()
	if err != nil {
		log.Println(fmt.Sprintf("ERROR: can't load cfg {%s}", cfg))
	}

	snsSvc, err := svc.NewSnsService(cfg.SnsArn, cfg.AwsRegion)
	if err != nil {
		log.Println(fmt.Sprintf("ERROR: can't init sns service {%s}", err))
	}

	svcManager := res.NewApiManager(log.Default(), snsSvc)

	engine := gin.New()
	engine.POST("/data-to-manage", svcManager.Publish)

	err = engine.Run()
	if err != nil {
		log.Println(fmt.Sprintf("ERROR: can't run gin engine {%s}", err))
	}
}
