package restaurantbiz

import (
	"context"
	"errors"

	"simple_rest_api.com/m/module/restaurant/restaurantmodel"
)

type CreateRestaurant interface {
	Create(ctx context.Context, data *restaurantmodel.RestaurantCreate) error
}

type CreateRestaurantBiz struct {
	store CreateRestaurant
}

func NewCreateRestaurantBiz(store CreateRestaurant) *CreateRestaurantBiz {
	return &CreateRestaurantBiz{store: store}
}

func (b *CreateRestaurantBiz) Create(ctx context.Context, data *restaurantmodel.RestaurantCreate) error {
	if *data.Name == "" {
		return errors.New("name can't be blank")
	}

	err := b.store.Create(ctx, data)

	return err
}
