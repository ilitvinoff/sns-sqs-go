package service

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	parquetS3 "github.com/xitongsys/parquet-go-source/s3"
	"github.com/xitongsys/parquet-go/writer"
	"sns-sqs/common/models"
	"time"
)

type S3ParquetHandler interface {
	WriteToKey(ctx context.Context, msg *models.Message) error
}

type S3ParquetService struct {
	Bucket  string
	Configs []*aws.Config
}

func NewS3ParquetService(bucket string, configs ...*aws.Config) *S3ParquetService {
	return &S3ParquetService{bucket, configs}
}

func (s *S3ParquetService) WriteToKey(ctx context.Context, msg *models.Message) error {
	fw, err := parquetS3.NewS3FileWriter(
		ctx,
		s.Bucket,
		fmt.Sprintf("%d", time.Now().Unix()),
		"",
		[]func(*s3manager.Uploader){}, s.Configs...)
	if err != nil {
		return err
	}
	defer fw.Close()

	pw, err := writer.NewJSONWriter(models.MessageSchema, fw, 4)
	if err != nil {
		return err
	}

	msgJson, err := msg.GetJSON()
	if err != nil {
		return err
	}

	err = pw.Write(msgJson)
	if err != nil {
		return err
	}

	_ = pw.WriteStop()

	return nil
}
