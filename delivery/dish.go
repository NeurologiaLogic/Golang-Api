package delivery

import (
	"github.com/gin-gonic/gin"
	"rest-api/services"
	"net/http"
)
type Delivery interface {
	Mount(group *gin.RouterGroup)
}

type delivery struct {
	dishService services.DishServices
}

func NewDishDelivery(dishServices services.DishServices) *delivery{
	return &delivery{dishServices}
}

func (d *delivery) Mount(group *gin.RouterGroup) {
	group.GET("/", d.getAllDish)
	group.POST("/", d.createDish)
	group.GET("/:id", d.getDishByID)
	group.PUT("/:id", d.UpdateDish)
	group.DELETE("/:id", d.DeleteDish)
}

func (d *delivery) getAllDish(c *gin.Context){
	// c.JSON(http.StatusOK,gin.H{
	// 	"status":"Success",
	// 	"message":[]string{"Dish1","Dish2","Dish3"},
	// })
	d.dishService.GetAllDish()
}
func (d *delivery) getDishByID(c *gin.Context){
	// c.JSON(http.StatusOK,gin.H{
	// 	"status":"Success",
	// 	"message":"Dish1",
	// })
	d.dishService.GetDishByID(1)
}
func (d *delivery) createDish(c *gin.Context){
	c.JSON(http.StatusOK,gin.H{
		"status":"Success",
		"message":"New Dish Created",
	})
}
func (d *delivery) UpdateDish(c *gin.Context){
	c.JSON(http.StatusOK,gin.H{
		"status":"Success",
		"message":"Dish Successfully Updated",
	})
}
func (d *delivery) DeleteDish(c *gin.Context){
	c.JSON(http.StatusOK,gin.H{
		"status":"Success",
		"message":"Dish Successfully Deleted",
	})
}
// dah bener