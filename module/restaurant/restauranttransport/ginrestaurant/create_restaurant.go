package ginrestaurant

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"simple_rest_api.com/m/module/restaurant/restaturantstorage"
	"simple_rest_api.com/m/module/restaurant/restaurantbiz"
	"simple_rest_api.com/m/module/restaurant/restaurantmodel"
)

func CreateRestaurant(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var data restaurantmodel.RestaurantCreate

		if err := c.ShouldBind(&data); err != nil {
			// bản chất gin.H là map[string]interface{}
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		store := restaturantstorage.NewSqlStore(db)
		biz := restaurantbiz.NewCreateRestaurantBiz(store)

		if err := biz.Create(c.Request.Context(), &data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, data)
	}
}
