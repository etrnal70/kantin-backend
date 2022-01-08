package model

type ImagePath struct {
	ID   uint   `json:"id"`
	Path string `json:"image_path"`
}

type MenuItem struct {
	ID          uint64    `json:"id"`
	Image       ImagePath `json:"image"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Price       int       `json:"price"`
}
