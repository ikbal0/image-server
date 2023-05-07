package entities

type Image struct {
	GormModel
	Name     string `json:"name"`
	ImageUrl string `json:"imageUrl"`
	UserID   uint   `json:"userId"`
}
