package component

import (
	"gorm.io/gorm"
	"simple_rest_api.com/m/component/uploadprovider"
)

type AppContext interface {
	GetMainDBConnection() *gorm.DB
	UploadProvider() uploadprovider.UploadProvider
}

type appCtx struct {
	db         *gorm.DB
	upProvider uploadprovider.UploadProvider
}

func NewAppContext(db *gorm.DB, upProvider uploadprovider.UploadProvider) *appCtx {
	return &appCtx{db: db, upProvider: upProvider}
}

func (c *appCtx) GetMainDBConnection() *gorm.DB {
	return c.db
}

func (ctx *appCtx) UploadProvider() uploadprovider.UploadProvider {
	return ctx.upProvider
}
