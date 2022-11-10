package model

type Dish struct {
	ID          int    `gorm:"primary_key"`
	Name        string `json:"name"`
	Image       string `json:"image"`
	Origin      string `json:"origin"`
	Type        string `json:"type"`
	Description string `json:"description"`
}