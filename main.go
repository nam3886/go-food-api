package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"simple_rest_api.com/m/component"
	"simple_rest_api.com/m/module/restaurant/restauranttransport/ginrestaurant"
)

func main() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalln("Error loading .env file")
	}

	db, err := connectDB()

	if err != nil {
		log.Fatalln(err)
	}

	if err := runService(db); err != nil {
		log.Fatalln(err)
	}
}

func connectDB() (*gorm.DB, error) {
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	name := os.Getenv("DB_DATABASE")
	username := os.Getenv("DB_USERNAME")
	password := os.Getenv("DB_PASSWORD")
	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", username, password, host, port, name)
	db, err := gorm.Open(mysql.Open(connectionString), &gorm.Config{})

	return db, err
}

func runService(db *gorm.DB) error {
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "pong"})
	})

	appCtx := component.NewAppContext(db)

	restaurants := r.Group("/restaurants")
	restaurants.POST("", ginrestaurant.CreateRestaurant(appCtx))
	restaurants.GET("", ginrestaurant.ListRestaurant(appCtx))
	restaurants.GET("/:id", ginrestaurant.GetRestaurant(appCtx))
	restaurants.PUT("/:id", ginrestaurant.UpdateRestaurant(appCtx))
	restaurants.DELETE("/:id", ginrestaurant.DeleteRestaurant(appCtx))

	return r.Run(fmt.Sprintf(":%s", os.Getenv("APP_PORT")))
}
