package res

import (
	"context"
	"encoding/json"
	"fmt"
	"sns-sqs/common/models"
)

func (m *ApiManager) HandleMessage() {
	output, err := m.SqsService.ReceiveMessage()
	if err != nil {
		m.Logger.Println(fmt.Sprintf("ERROR: can't receive message {%s}", err))
		return
	}

	for _, unit := range output.Messages {
		msg := &models.Message{}
		err = json.Unmarshal([]byte(*unit.Body), &msg)
		if err != nil {
			m.Logger.Println(fmt.Sprintf("ERROR: сannot retrieve the message from the output body {%s}", err))
			continue
		}

		isMsgExists, err := m.DynamoDbService.IsItemExists(msg.ID)
		if err != nil {
			m.Logger.Println(fmt.Sprintf("ERROR: сannot define if message already exists {%s}", err))
			continue
		}

		if !isMsgExists {
			err = m.DynamoDbService.PutMessage(msg.ID)
			if err != nil {
				m.Logger.Println(fmt.Sprintf("ERROR: сannot put message to ddb {%s}", err))
				continue
			}

			m.S3ParquetService.WriteToKey(context.Background(), msg)
		}

		_, err = m.SqsService.DeleteMessage(unit.ReceiptHandle)
		if err != nil {
			m.Logger.Println(fmt.Sprintf("ERROR: сannot delete message from sqs {%s}", err))
		}
	}

}
