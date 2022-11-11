package services
import (
	"rest-api/model"
	"rest-api/repository"
	"github.com/gin-gonic/gin"
	validate "github.com/go-playground/validator/v10"
	"strconv"
)
type DishServices interface{
	GetAllDish() ([]model.Dish, error)
	GetDishByID(c *gin.Context) (model.Dish, error)
	CreateDish(c *gin.Context) error
	UpdateDish(c *gin.Context) error
	DeleteDish(c *gin.Context) error
}

type dishServices struct {
	repo repository.DishRepository
}

func NewDishServices(repo repository.DishRepository)DishServices{
	return &dishServices{repo}
}

//getting accessable methods through newServices
//format func method nama struct diikuti dengan method dri interface
func (s *dishServices) GetAllDish() ([]model.Dish, error){
	return s.repo.GetAllDish()
}

func (s *dishServices) GetDishByID(c *gin.Context) (model.Dish, error){
	id,err:=strconv.Atoi(c.Param("id"))
	if err!=nil{
		return model.Dish{},err
	}
	return s.repo.DeleteDish(id)
}

func (s *dishServices) CreateDish(c *gin.Context) (error){
	dish:=model.Dish{}
	if err:=c.Bind(&dish);err!=nil{
		return err
	}
	validate:=validate.New()
	if err:=validate.Struct(dish);err!=nil{
		return err
	}
	return s.repo.CreateDish(&dish)
}
func (s *dishServices) UpdateDish(c *gin.Context) (error){
	return nil
}
func (s *dishServices) DeleteDish(c *gin.Context) (error){
	id,err:=strconv.Atoi(c.Param("id"))
	if err!=nil{
		return err
	}
	return s.repo.DeleteDish(id)
}

