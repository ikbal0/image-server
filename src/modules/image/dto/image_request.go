package dto

type ImageRequestBody struct {
	Name     string `json:"name"`
	ImageUrl string `json:"imageUrl"`
	UserID   uint   `json:"userId"`
}
