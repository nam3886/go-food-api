package restaurantmodel

import (
	"errors"
	"strings"

	"simple_rest_api.com/m/common"
)

// business model
type Restaurant struct {
	common.SQLModel `json:",inline"`
	Slug            string `json:"slug"`
	Name            string `json:"name" gorm:"column:name"` // chỉ định tên cột
}

// đặt tên cho table, nếu ko có sẽ lấy mặc định như bên model laravel
func (Restaurant) TableName() string {
	return "features"
}

type RestaurantCreate struct {
	common.SQLModel `json:",inline"`
	Slug            string  `json:"slug"`
	Name            *string `json:"name" gorm:"column:name"`
}

func (RestaurantCreate) TableName() string {
	return "features"
}

func (r *RestaurantCreate) Validate() error {
	name := strings.TrimSpace(*r.Name)

	if len(name) == 0 {
		return errors.New("restaurant name cannot be blank")
	}

	return nil
}

type RestaurantUpdate struct {
	common.SQLModel `json:",inline"`
	Slug            string  `json:"slug"`
	Name            *string `json:"name" gorm:"column:name"`
}

func (RestaurantUpdate) TableName() string {
	return "features"
}

func (r *RestaurantUpdate) Validate() error {
	name := strings.TrimSpace(*r.Name)

	if len(name) == 0 {
		return errors.New("restaurant name cannot be blank")
	}

	return nil
}
