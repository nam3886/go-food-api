package restaurantstorage

import (
	"context"

	"simple_rest_api.com/m/common"
	"simple_rest_api.com/m/modules/restaurant/restaurantmodel"
)

func (s *sqlStore) Delete(ctx context.Context, id int) error {
	db := s.db

	if err := db.Where("id = ?", id).Delete(&restaurantmodel.Restaurant{}).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}
