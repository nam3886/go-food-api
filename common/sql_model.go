package common

import (
	"time"

	"gorm.io/gorm"
)

type SQLModel struct {
	// quy đổi ra json sẽ có key là id (tiện lợi cho việc muốn alias lại key), omitempty nếu giá trị là zero value => remove key id luôn
	Id int `json:"id,omitempty" gorm:"primaryKey`
	// có con trỏ => check zero value = null nếu ko có check zero = false
	Active    *bool          `json:"active,omitempty" gorm:"default:1"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at,omitempty"`
}
