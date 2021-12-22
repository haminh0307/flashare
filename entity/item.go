package entity

type Item struct {
	ID          string `bson:"_id,omitempty" json:"id"`
	Title       string `bson:"title" json:"title"`
	Category    string `bson:"category" json:"category"`
	PhotosLink  string `bson:"photos_link" json:"photos_link"`
	Description string `bson:"description" json:"description"`
	DueDate     string `bson:"due_date" json:"due_date"`
	UploadBy    string `bson:"upload_by" json:"upload_by"`
	Status      string `bson:"status" json:"status"`
}
