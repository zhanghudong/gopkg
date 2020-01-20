package appctx

import (
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"github.com/jinzhu/gorm"
)

type Context struct {
	*gin.Context
	Mysql *gorm.DB
	Redis *redis.Client
}

func NewContext(ctx *gin.Context, mysql *gorm.DB, redis *redis.Client) *Context {
	return &Context{
		Context: ctx,
		Mysql:   mysql,
		Redis:   redis,
	}
}
