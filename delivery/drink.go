package delivery

import (
	"github.com/gin-gonic/gin"
	"rest-api/services"
	"net/http"
)
type DrinkDelivery interface {
	Mount(group *gin.RouterGroup)
}

type drinkDelivery struct {
	drinkServices services.DrinkServices
}

func NewDrinkDelivery(drinkServices services.DrinkServices) DrinkDelivery{
	return &drinkDelivery{drinkServices}
}

func (d *drinkDelivery) Mount(group *gin.RouterGroup) {
	group.GET("/", d.getAllDrink)
	// group.POST("/", d.createDish)
	// group.GET("/:id", d.getDishByID)
	// group.PUT("/:id", d.UpdateDish)
	// group.DELETE("/:id", d.DeleteDish)
}

func (d *drinkDelivery) getAllDrink(c *gin.Context){
	// c.JSON(http.StatusOK,gin.H{
	// 	"status":"Success",
	// 	"message":[]string{"Dish1","Dish2","Dish3"},
	// })
	data, err := d.drinkServices.GetAllDrink()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": "Error",
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": "Success",
		"message": data,
	})
}
// func (d *delivery) getDishByID(c *gin.Context){
// 	c.JSON(http.StatusOK,gin.H{
// 		"status":"Success",
// 		"message":"Dish1",
// 	})
// 	d.dishService.GetDishByID(1)
// }
// func (d *delivery) createDish(c *gin.Context){
// 	c.JSON(http.StatusOK,gin.H{
// 		"status":"Success",
// 		"message":"New Dish Created",
// 	})
// }
// func (d *delivery) UpdateDish(c *gin.Context){
// 	c.JSON(http.StatusOK,gin.H{
// 		"status":"Success",
// 		"message":"Dish Successfully Updated",
// 	})
// }
// func (d *delivery) DeleteDish(c *gin.Context){
// 	c.JSON(http.StatusOK,gin.H{
// 		"status":"Success",
// 		"message":"Dish Successfully Deleted",
// 	})
// }
// dah bener