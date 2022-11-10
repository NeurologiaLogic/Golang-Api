package services

import (
	"rest-api/model"
	"rest-api/repository"
)
type service interface{
	GetAllDish() ([]*model.Dish, error)
	GetDishByID(id int) (*model.Dish, error)
	CreateDish(dish *model.Dish) error
	UpdateDish(dish *model.Dish) error
	DeleteDish(id int) error
}

type DishServices struct {
	repo repository.Repository
}

func NewDishServices(repo repository.Repository)*DishServices{
	return &DishServices{repo}
}

//getting accessable methods through newServices

//format func method nama struct diikuti dengan method dri interface
func (s *DishServices) GetAllDish() ([]model.Dish, error){
	return s.repo.GetAllDish()
}

func (s *DishServices) GetDishByID(id int) (model.Dish, error){
	return s.repo.GetDishByID(id)
}
func (s *DishServices) CreateDish(dish *model.Dish) (error){
	return nil
}
func (s *DishServices) UpdateDish(dish *model.Dish) (error){
	return nil
}
func (s *DishServices) DeleteDish(id int) (error){
	return nil
}

