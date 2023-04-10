package models

type Image struct {
	GormModel
	Title    string
	Caption  string
	ImageUrl string
	UserID   uint
}
