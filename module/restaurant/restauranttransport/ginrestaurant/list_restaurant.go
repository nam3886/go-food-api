package ginrestaurant

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"simple_rest_api.com/m/common"
	"simple_rest_api.com/m/component"
	"simple_rest_api.com/m/module/restaurant/restaturantstorage"
	"simple_rest_api.com/m/module/restaurant/restaurantbiz"
	"simple_rest_api.com/m/module/restaurant/restaurantmodel"
)

func ListRestaurant(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var filter restaurantmodel.Filter

		if err := c.ShouldBind(&filter); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		var paging common.Paging

		if err := c.ShouldBind(&paging); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		paging.Fulfill()

		store := restaturantstorage.NewSqlStore(appCtx.GetMainDBConnection())
		biz := restaurantbiz.NewListRestaurantBiz(store)
		result, err := biz.ListRestaurant(c.Request.Context(), &filter, &paging)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, common.NewSuccessResponse(result, paging, filter))
	}
}
