/*
使用 Gin 框架 + MySQL 实现 QQ 的好友功能
加好友
删好友
查看所有好友
好友分组
好友搜索
*/
package main

import (
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"lesson4/lesson4/api"
)

func main() {
	api.InitRouter()

}
