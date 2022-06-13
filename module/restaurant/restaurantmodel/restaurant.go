package restaurantmodel

import (
	"errors"
	"strings"
	"time"
)

// business model
type Restaurant struct {
	Id        int        `json:"id,omitempty" gorm:"primaryKey"` // quy đổi ra json sẽ có key là id (tiện lợi cho việc muốn alias lại key), omitempty nếu giá trị là zero value => remove key id luôn
	Slug      string     `json:"slug"`
	Name      string     `json:"name" gorm:"column:name"`           // chỉ định tên cột
	Active    *bool      `json:"active,omitempty" gorm:"default:1"` // có con trỏ => check zero value = null nếu ko có check zero = false
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at,omitempty" gorm:"default:NULL"`
}

// đặt tên cho table, nếu ko có sẽ lấy mặc định như bên model laravel
func (Restaurant) TableName() string {
	return "features"
}

type RestaurantCreate struct {
	Id     int     `json:"id,omitempty" gorm:"primaryKey"`
	Slug   string  `json:"slug"`
	Name   *string `json:"name" gorm:"column:name"`
	Active *bool   `json:"active,omitempty" gorm:"default:1"`
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
	Id     int     `json:"id,omitempty" gorm:"primaryKey"`
	Slug   string  `json:"slug"`
	Name   *string `json:"name" gorm:"column:name"`
	Active *bool   `json:"active,omitempty" gorm:"default:1"`
}

func (RestaurantUpdate) TableName() string {
	return "features"
}
