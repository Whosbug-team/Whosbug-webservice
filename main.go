package main

import (
	. "webService_Refactoring/modules"
	. "webService_Refactoring/routes"
)

func main() {
	// 引用数据库
	InitDB()

	// 启动路由
	InitRouter()
}
