package main

import (
	"fmt"
	"lastwork/api"
	"lastwork/initity"
)

func main() {
	err := initity.IniMysql()
	if err != nil {
		fmt.Println("mysql数据库连接失败")
		fmt.Println(err.Error())
		return
	}
	//err = initity.IniRedis()
	//if err != nil {
	//	fmt.Println("redis数据库连接失败")
	//}
	api.InitRouter()
}
