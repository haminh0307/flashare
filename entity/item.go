package entity

type Item struct {
	ID          string `json:"msg_id"`
	Title       string `json:"title"`
	Category    string `json:"category"`
	PhotosLink  string `json:"photos_link"`
	Description string `json:"description"`
	DueDate     string `json:"due_date"`
	UploadBy    string `json:"upload_by"`
	Status      string `json:"status"`
}
