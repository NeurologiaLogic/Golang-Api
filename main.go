package main

import (
	"fmt"
	"log"
	"os"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"rest-api/delivery"
	"rest-api/repository"
	"rest-api/services"
)


func setupRouter() *gin.Engine{
	db:=setupDB()
	r:=gin.Default()
	r.TrustedPlatform = gin.PlatformGoogleAppEngine
	// dishRepo:=repository.NewDishTestRepository(db)
	dishRepo:=repository.NewDishRepository(db)
	dishServices:=services.NewDishServices(dishRepo)
	dishDelivery:=delivery.NewDishDelivery(dishServices)
	dishGroup := r.Group("/dish")
	dishDelivery.Mount(dishGroup)


	drinkRepo :=repository.NewDrinkRepository(db)
	drinkServices := services.NewDrinkServices(drinkRepo)
	drinkDelivery := delivery.NewDrinkDelivery(drinkServices)
	drinkGroup := r.Group("/drink")
	drinkDelivery.Mount(drinkGroup)
	return r
}

func setupDB() *gorm.DB{
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	s3Bucket := os.Getenv("FILENAME")
	dsn := fmt.Sprintf("root:@tcp(127.0.0.1:3306)/%s?charset=utf8mb4&parseTime=True&loc=Local",s3Bucket)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
  if err != nil {
    panic("failed to connect database")
  }
	return db
}
func main(){
	r:=setupRouter()
	r.Run("localhost:1234")
}