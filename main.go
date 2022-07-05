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
	"simple_rest_api.com/m/component/uploadprovider"
	"simple_rest_api.com/m/middleware"
	"simple_rest_api.com/m/modules/restaurant/restauranttransport/ginrestaurant"
	"simple_rest_api.com/m/modules/upload/uploadtransport/ginupload"
)

func main() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalln("Error loading .env file")
	}

	s3Provider := getS3Provider()
	db, err := connectDB()

	if err != nil {
		log.Fatalln(err)
	}

	if err := runService(db, s3Provider); err != nil {
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

func getS3Provider() uploadprovider.UploadProvider {
	s3BucketName := os.Getenv("S3_BUCKET_NAME")
	s3Region := os.Getenv("S3_REGION")
	s3APIKey := os.Getenv("S3_API_KEY")
	s3SecretKey := os.Getenv("S3_SECRET_KEY")
	s3Domain := os.Getenv("S3_DOMAIN")
	s3Provider := uploadprovider.NewS3Provider(s3BucketName, s3Region, s3APIKey, s3SecretKey, s3Domain)

	return s3Provider
}

func runService(db *gorm.DB, upProvider uploadprovider.UploadProvider) error {
	appCtx := component.NewAppContext(db, upProvider)
	r := gin.Default()

	r.Use(middleware.Recover(appCtx))

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "pong"})
	})

	r.POST("upload", ginupload.Upload(appCtx))

	restaurants := r.Group("/restaurants")
	restaurants.POST("", ginrestaurant.CreateRestaurant(appCtx))
	restaurants.GET("", ginrestaurant.ListRestaurant(appCtx))
	restaurants.GET("/:id", ginrestaurant.GetRestaurant(appCtx))
	restaurants.PUT("/:id", ginrestaurant.UpdateRestaurant(appCtx))
	restaurants.DELETE("/:id", ginrestaurant.DeleteRestaurant(appCtx))

	return r.Run(fmt.Sprintf(":%s", os.Getenv("APP_PORT")))
}
