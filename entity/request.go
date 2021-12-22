package entity

type Request struct {
	ID       string `bson:"_id,omitempty" json:"id"`
	Item     string `bson:"item" json:"item"`
	Sender   string `bson:"sender" json:"sender"`
	Receiver string `bson:"receiver" json:"receiver"`
	Status   string `bson:"status" json:"status"`
}
