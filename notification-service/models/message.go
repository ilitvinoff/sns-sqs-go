package models

import "fmt"

type Message struct {
	ID  string `json:"id" binding:"required"`
	Msg string `json:"msg" binding:"required"`
}

func (m Message) String() string {
	return fmt.Sprintf("ID: %s; Msg: %s", m.ID, m.Msg)
}
