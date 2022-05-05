package main

import (
	"todomvc-app-template-golang/db"
	"todomvc-app-template-golang/routers"
)

func main() {
	db.InitMysql() //创建数据库//连接数据库

	routers.SetRouters() //注册路由
}
