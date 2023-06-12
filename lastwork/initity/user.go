package initity

import (
	"context"
	"github.com/go-redis/redis/v8"
	"github.com/jinzhu/gorm"
	"lastwork/global"
	"lastwork/model"
)

func IniMysql() error {
	dsn := "root:369248@tcp(127.0.0.1:3306)/db3?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	global.DB, err = gorm.Open("mysql", dsn)
	global.DB.AutoMigrate(&model.User{}, &model.Problem{}, &model.Answer{})
	return err
}
func IniRedis() (err error) {
	global.Rbq = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	_, err = global.Rbq.Ping(context.Background()).Result()
	return err
}
