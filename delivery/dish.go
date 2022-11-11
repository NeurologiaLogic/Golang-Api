package delivery

import (
	"net/http"
	"rest-api/services"

	"github.com/gin-gonic/gin"
)
type DishDelivery interface {
	Mount(group *gin.RouterGroup)
}

type dishDelivery struct {
	dishService services.DishServices
}

func NewDishDelivery(dishServices services.DishServices) DishDelivery{
	return &dishDelivery{dishServices}
}

func (d *dishDelivery) Mount(group *gin.RouterGroup) {
	group.GET("/", d.getAllDish)
	group.POST("/", d.createDish)
	group.GET("/:id", d.getDishByID)
	group.PUT("/:id", d.UpdateDish)
	group.DELETE("/:id", d.DeleteDish)
}

func (d *dishDelivery) getAllDish(c *gin.Context){
	res,err:=d.dishService.GetAllDish()
	if err != nil{
		c.JSON(http.StatusInternalServerError,gin.H{
			"status":"Error",
			"message":err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK,gin.H{
		"status":"Success",
		"message":res,
	})
}
func (d *dishDelivery) getDishByID(c *gin.Context){
	res,err:=d.dishService.GetAllDish()
	if err != nil{
		c.JSON(http.StatusInternalServerError,gin.H{
			"status":"Error",
			"message":err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK,gin.H{
		"status":"Success",
		"message":res,
	})
}
func (d *dishDelivery) createDish(c *gin.Context){
		err:=d.dishService.CreateDish(c)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status": "Error",
				"message": err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK,gin.H{
			"status":"Success",
			"message":"New Dish Created",
		})
}
func (d *dishDelivery) UpdateDish(c *gin.Context){
	c.JSON(http.StatusOK,gin.H{
		"status":"Success",
		"message":"Dish Successfully Updated",
	})
}
func (d *dishDelivery) DeleteDish(c *gin.Context){
	err:=d.dishService.DeleteDish(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "Error",
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK,gin.H{
		"status":"Success",
		"message":"Dish Deleted",
	})
}
// dah bener