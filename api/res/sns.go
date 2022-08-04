package res

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"sns-sqs/common/models"
)

func (m *ApiManager) Publish(c *gin.Context) {
	msgModel := models.Message{}

	err := c.BindJSON(&msgModel)
	if err != nil {
		m.Logger.Println(fmt.Sprintf("ERROR: can't bind request to publish into sns {%s}", err))
		return
	}

	msgString, err := json.Marshal(msgModel)
	if err != nil {
		m.Logger.Println(fmt.Sprintf("ERROR: can't marshal message {%s}", err))
		return
	}

	_, err = m.SnsService.Publish(string(msgString))
	if err != nil {
		m.Logger.Println(fmt.Sprintf("ERROR: can't publish message {%s}", err))
		return
	}

}
