package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"simple_rest_api.com/m/module/restaurant/restaurantmodel"
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

	restaurants := r.Group("/restaurants")
	restaurants.POST("", ginrestaurant.CreateRestaurant(db))
	restaurants.GET("", func(c *gin.Context) {
		var data []restaurantmodel.Restaurant

		type Filter struct {
			CityId int `json:"city_id" form:"city_id"` // tag form cần khi muốn nhận giá trị từ query string với key = city_id
		}

		var filter Filter

		if err := c.ShouldBind(&filter); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		fmt.Println(filter)
		newDb := db.Debug()

		if filter.CityId > 0 {
			newDb = newDb.Where("city_id = ?", filter.CityId)
		}

		if err := newDb.Find(&data).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, data)
	})
	restaurants.GET("/:id", func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		var data restaurantmodel.Restaurant

		if err := db.Where("id = ?", id).First(&data).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, data)
	})
	restaurants.PUT("/:id", func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		var data restaurantmodel.Restaurant

		if err := c.ShouldBind(&data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if err := db.Where("id = ?", id).Updates(&data).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, data)
	})
	restaurants.DELETE("/:id", func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if err := db.Where("id = ?", id).Delete(&restaurantmodel.Restaurant{}).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"data": "ok"})
	})

	return r.Run(fmt.Sprintf(":%s", os.Getenv("APP_PORT")))
}