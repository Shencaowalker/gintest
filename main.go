package main

import (
	"gintest1/common"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	//"github.com/spf13/viper"
)

func main() {
	//InitConfig()
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

//下面代码本意读取config目录下的yml   现在不生效
/*
func InitConfig() {
	v := viper.New()
	v.SetConfigName("application")
	v.SetConfigType("yml")
	v.AddConfigPath("./config")
	err := v.ReadInConfig()
	if err != nil {
		return
	}
}
*/
