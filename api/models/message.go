package models

type Message struct {
	ID  string `json:"id" binding:"required"`
	Msg string `json:"msg" binding:"required"`
}
