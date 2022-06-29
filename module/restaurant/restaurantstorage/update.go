package restaurantstorage

import (
	"context"

	"simple_rest_api.com/m/common"
	"simple_rest_api.com/m/module/restaurant/restaurantmodel"
)

func (s *sqlStore) Update(ctx context.Context, id int, data *restaurantmodel.RestaurantUpdate) error {
	db := s.db

	if err := db.Where("id = ?", id).Updates(data).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}
