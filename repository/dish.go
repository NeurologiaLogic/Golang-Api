package repository

import (
	"rest-api/model"

	"gorm.io/gorm"
)

type DishRepository interface {
	GetAllDish() ([]model.Dish, error)
	GetDishByID(id int) (model.Dish, error)
	CreateDish(dish *model.Dish) error
	UpdateDish(dish *model.Dish) error
	DeleteDish(id int) error
}

type dishRepository struct {
	db *gorm.DB
}

//entry point
func NewDishRepository(db *gorm.DB) DishRepository{
  db.AutoMigrate(&model.Dish{})
	return &dishRepository{db}
}

func (r *dishRepository) GetAllDish() ([]model.Dish, error){
	var dishes []model.Dish
	if 	err:=r.db.Find(&dishes).Error;err!= nil {
		return nil, err
	}
	return dishes,nil
}
func (r *dishRepository) GetDishByID(id int) (model.Dish, error){
	var dish model.Dish
	if err:=r.db.Find(&dish, id).Error;err != nil {
		return dish, err
	}
	return dish,nil
}
func (r *dishRepository) CreateDish(dish *model.Dish) (error){
	createdDish:=r.db.Create(&dish)
	if createdDish.Error != nil {
		return createdDish.Error
	}
	return nil
}
func (r *dishRepository) UpdateDish(dish *model.Dish) (error){
	return nil
}
func (r *dishRepository) DeleteDish(id int) (error){
	var dish model.Dish
	if err:=r.db.Delete(&dish, id).Error;err != nil {
		return err
	}
	return nil
}
