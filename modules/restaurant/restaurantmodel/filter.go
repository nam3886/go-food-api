package restaurantmodel

type Filter struct {
	CityId int `json:"city_id,omitempty" form:"city_id"` // tag form cần khi muốn nhận giá trị từ query string với key = city_id
}
