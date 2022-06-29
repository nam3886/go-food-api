package restaurantmodel

import (
	"errors"
	"strings"

	"simple_rest_api.com/m/common"
)

const EntityName = "Restaurant"

// business model
type Restaurant struct {
	common.SQLModel `json:",inline"`
	Slug            string         `json:"slug"`
	Name            string         `json:"name" gorm:"column:name"` // chỉ định tên cột
	CityId          int            `json:"city_id,omitempty" gorm:"default:NULL"`
	Logo            *common.Image  `json:"logo" gorm:"default:NULL"`
	Cover           *common.Images `json:"cover" gorm:"default:NULL"`
}

// đặt tên cho table, nếu ko có sẽ lấy mặc định như bên model laravel
func (Restaurant) TableName() string {
	return "restaurants"
}

type RestaurantCreate struct {
	Active *bool          `json:"active,omitempty" gorm:"default:1"`
	Slug   string         `json:"slug"`
	Name   *string        `json:"name" gorm:"column:name"`
	CityId int            `json:"city_id,omitempty" gorm:"default:NULL"`
	Logo   *common.Image  `json:"logo" gorm:"default:NULL"`
	Cover  *common.Images `json:"cover" gorm:"default:NULL"`
}

func (RestaurantCreate) TableName() string {
	return "restaurants"
}

func (r *RestaurantCreate) Validate() error {
	name := strings.TrimSpace(*r.Name)

	if len(name) == 0 {
		return errors.New("restaurant name cannot be blank")
	}

	return nil
}

type RestaurantUpdate struct {
	Active *bool          `json:"active,omitempty" gorm:"default:1"`
	Slug   string         `json:"slug"`
	Name   *string        `json:"name" gorm:"column:name"`
	CityId int            `json:"city_id,omitempty" gorm:"default:NULL"`
	Logo   *common.Image  `json:"logo" gorm:"default:NULL"`
	Cover  *common.Images `json:"cover" gorm:"default:NULL"`
}

func (RestaurantUpdate) TableName() string {
	return "restaurants"
}

func (r *RestaurantUpdate) Validate() error {
	name := strings.TrimSpace(*r.Name)

	if len(name) == 0 {
		return errors.New("restaurant name cannot be blank")
	}

	return nil
}
