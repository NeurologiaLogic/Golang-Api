package repository

import(
	"gorm.io/gorm"
	"rest-api/model"
)

type Repository interface {
	GetAllDish() ([]model.Dish, error)
	GetDishByID(id int) (model.Dish, error)
	CreateDish(dish *model.Dish) error
	UpdateDish(dish *model.Dish) error
	DeleteDish(id int) error
}

type repository struct {
	db *gorm.DB
}

//entry point
func NewDishRepository(db *gorm.DB) *repository{
	return &repository{db}
}

func (r *repository) GetAllDish() ([]model.Dish, error){
	var dish []model.Dish
	return dish,nil
}
func (r *repository) GetDishByID(id int) (model.Dish, error){
	var dish model.Dish
	return dish,nil
}
func (r *repository) CreateDish(dish *model.Dish) (error){
	return nil
}
func (r *repository) UpdateDish(dish *model.Dish) (error){
	return nil
}
func (r *repository) DeleteDish(id int) (error){
	return nil
}
