package models

import (
	"encoding/json"
	"fmt"
)

type Message struct {
	ID  string `json:"id" binding:"required"`
	Msg string `json:"msg" binding:"required"`
}

var MessageSchema = `{
        "Tag":"name=parquet-go-root",
        "Fields":[
		    {"Tag":"name=id, type=string, convertedtype=UTF8, repetitiontype=OPTIONAL"},
		    {"Tag":"name=msg, type=string", convertedtype=UTF8, repetitiontype=OPTIONAL"},
}`

func (m Message) String() string {
	return fmt.Sprintf("ID: %s; Msg: %s", m.ID, m.Msg)
}

func (m *Message) GetJSON() (string, error) {
	res, err := json.Marshal(m)
	if err != nil {
		return "", err
	}
	return string(res), nil
}
