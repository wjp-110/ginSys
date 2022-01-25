package main

import (
	"ginSys/common"
	"ginSys/routes"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
	"os"
)

func main() {
	//读取配置文件
	InitConfig()

	//初始化db
	db := common.InitDB()
	defer db.Close()

	r := gin.Default()

	//加载路由
	r = routes.CollectRoute(r)

	panic(r.Run())
}

func InitConfig()  {
	workDir, _ := os.Getwd()
	viper.SetConfigName("application")
	viper.SetConfigType("yml")
	viper.AddConfigPath(workDir+ "/config")
	err:= viper.ReadInConfig()
	if err != nil {
		panic("read config is error->%v"+ err.Error())
	}
}
