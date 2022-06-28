package restaurantbiz

import (
	"context"

	"simple_rest_api.com/m/module/restaurant/restaurantmodel"
)

type DeleteRestaurantStore interface {
	FindDataByCondition(
		ctx context.Context,
		conditions map[string]interface{},
		moreKeys ...string,
	) (*restaurantmodel.Restaurant, error)

	Delete(ctx context.Context, id int) error
}

type deleteRestaurantBiz struct {
	store DeleteRestaurantStore
}

func NewDeleteRestaurantBiz(store DeleteRestaurantStore) *deleteRestaurantBiz {
	return &deleteRestaurantBiz{store: store}
}

func (b *deleteRestaurantBiz) DeleteRestaurant(ctx context.Context, id int) error {
	_, err := b.store.FindDataByCondition(ctx, map[string]interface{}{"id": id})

	if err != nil {
		return err
	}

	err = b.store.Delete(ctx, id)

	if err != nil {
		return err
	}

	return err
}
