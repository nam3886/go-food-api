package restaurantbiz

import (
	"context"

	"simple_rest_api.com/m/common"
	"simple_rest_api.com/m/modules/restaurant/restaurantmodel"
)

type UpdateRestaurantStore interface {
	FindDataByCondition(
		ctx context.Context,
		conditions map[string]interface{},
		moreKeys ...string,
	) (*restaurantmodel.Restaurant, error)

	Update(ctx context.Context, id int, data *restaurantmodel.RestaurantUpdate) error
}

type updateRestaurantBiz struct {
	store UpdateRestaurantStore
}

func NewUpdateRestaurantBiz(store UpdateRestaurantStore) *updateRestaurantBiz {
	return &updateRestaurantBiz{store: store}
}

func (b *updateRestaurantBiz) UpdateRestaurant(ctx context.Context, id int, data *restaurantmodel.RestaurantUpdate) error {
	if err := data.Validate(); err != nil {
		return common.ErrCannotUpdateEntity(restaurantmodel.EntityName, err)
	}

	_, err := b.store.FindDataByCondition(ctx, map[string]interface{}{"id": id})

	if err != nil {
		if err != common.RecordNotFound {
			return common.ErrCannotGetEntity(restaurantmodel.EntityName, err)
		}

		return common.ErrCannotGetEntity(restaurantmodel.EntityName, err)
	}

	if err = b.store.Update(ctx, id, data); err != nil {
		return common.ErrCannotUpdateEntity(restaurantmodel.EntityName, err)
	}

	return nil
}
