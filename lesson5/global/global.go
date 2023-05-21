package global

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var DB *gorm.DB
var Rbq *redis.Client
var err error
var SecretSignKey = []byte("guoyiyang")

func Initmysql() {

	dsn := "root:369248@tcp(127.0.0.1:3306)/db2?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open("mysql", dsn)
	if err != nil {
		fmt.Println(err)
		fmt.Println("连接数据库失败")
		return
	}
}
func Initredis() (err error) {
	Rbq = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	_, err = Rbq.Ping(context.Background()).Result()
	if err != nil {
		return err
	}
	return nil
}
