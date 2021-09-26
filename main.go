package main

import (
	"gintest1/common"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	//初始化数据库 创建表
	db := common.InitDB()
	//完事儿后关闭数据库
	defer db.Close()
	//创建gin默认实例
	r := gin.Default()
	//加上路由选择
	r = CollectRoute(r)
	panic(r.Run()) // 监听并在 0.0.0.0:8080 上启动服务
}
