package repository

import (
	"rest-api/model"
	"gorm.io/gorm"
)

type DrinkRepository interface {
	GetAllDrink() ([]model.Drink, error)
	GetDrinkByID(id int) (model.Drink, error)
	CreateDrink(drink *model.Drink) error
	UpdateDrink(drink *model.Drink) error
	DeleteDrink(id int) error
}

type drinkRepository struct {
	db *gorm.DB
}

//entry point
func NewDrinkRepository(db *gorm.DB) DrinkRepository{
	db.AutoMigrate(&model.Drink{})
	return &drinkRepository{db}
}

func (r *drinkRepository) GetAllDrink() ([]model.Drink, error){
	var drink []model.Drink
	return drink,nil
}
func (r *drinkRepository) GetDrinkByID(id int) (model.Drink, error){
	var drink model.Drink
	return drink,nil
}
func (r *drinkRepository) CreateDrink(drink *model.Drink) (error){
	return nil
}
func (r *drinkRepository) UpdateDrink(drink *model.Drink) (error){
	return nil
}
func (r *drinkRepository) DeleteDrink(id int) (error){
	return nil
}
