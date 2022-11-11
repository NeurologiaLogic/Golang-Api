package services

import (
	"rest-api/model"
	"rest-api/repository"
)
type DrinkServices interface{
	GetAllDrink() ([]model.Drink, error)
	// GetDrinkByID(id int) (*model.Drink, error)
	// CreateDrink(Drink *model.Drink) error
	// UpdateDrink(Drink *model.Drink) error
	// DeleteDrink(id int) error
}

type drinkServices struct {
	repo repository.DrinkRepository
}

func NewDrinkServices(repo repository.DrinkRepository)DrinkServices{
	return &drinkServices{repo}
}

//getting accessable methods through newServices

//format func method nama struct diikuti dengan method dri interface
func (s *drinkServices) GetAllDrink() ([]model.Drink, error){
	return s.repo.GetAllDrink()
}

// func (s *drinkServices) GetDrinkByID(id int) (model.Drink, error){
// 	return s.repo.GetDrinkByID(id)
// }
// func (s *drinkServices) CreateDrink(Drink *model.Drink) (error){
// 	return nil
// }
// func (s *drinkServices) UpdateDrink(Drink *model.Drink) (error){
// 	return nil
// }
// func (s *drinkServices) DeleteDrink(id int) (error){
// 	return nil
// }

