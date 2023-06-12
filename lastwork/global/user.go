package global

import (
	"github.com/go-redis/redis/v8"
	"github.com/jinzhu/gorm"
)

var DB *gorm.DB
var Rbq *redis.Client
