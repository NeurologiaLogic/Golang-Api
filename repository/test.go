package repository

import(
	"gorm.io/gorm"
	"rest-api/model"
	"log"
)

type wwww interface {
	GetAllDish() ([]model.Dish, error)
	GetDishByID(id int) (model.Dish, error)
	CreateDish(dish *model.Dish) error
	UpdateDish(dish *model.Dish) error
	DeleteDish(id int) error
}

type repositoryTest struct {
	db *gorm.DB
}

//entry point
func NewDishTestRepository(db *gorm.DB) *repositoryTest{
	return &repositoryTest{db}
}

func (r *repositoryTest) GetAllDish() ([]model.Dish, error){
	// dish:=[]string{"rice","noodle","bread"}
	log.Println("get all")
	var dish []model.Dish
	return dish,nil
}
func (r *repositoryTest) GetDishByID(id int) (model.Dish, error){
	var dish model.Dish
	log.Println("get by id")
	return dish,nil
}
func (r *repositoryTest) CreateDish(dish *model.Dish) (error){
	return nil
}
func (r *repositoryTest) UpdateDish(dish *model.Dish) (error){
	return nil
}
func (r *repositoryTest) DeleteDish(id int) (error){
	return nil
}
