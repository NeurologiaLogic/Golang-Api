package model

type Dish struct {
	ID          int    `gorm:"primary_key"`
	Name        string `json:"name" xml:"name" validate:"required"`
	Image       string `json:"image" xml:"image" validate:"required"`
	Origin      string `json:"origin" xml:"origin" validate:"required"`
	Type        string `json:"type" xml:"type" validate:"required"`
	Description string `json:"description" xml:"description" validate:"required"`
}
type DishTemplate struct {
	ID          int    `gorm:"primary_key"`
	Name        string `json:"name" xml:"name" `
	Image       string `json:"image" xml:"image" `
	Origin      string `json:"origin" xml:"origin"`
	Type        string `json:"type" xml:"type"`
	Description string `json:"description" xml:"description"`
}