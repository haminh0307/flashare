package entity

type Message struct {
	ID       string `json:"id" binding:"required"`
	Sender   string `json:"sender" binding:"required"`
	Receiver string `json:"receiver" binding:"required"`
	Content  string `json:"content" binding:"required"`
	Time     string `json:"time" binding:"required"`
}
