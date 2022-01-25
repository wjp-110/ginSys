package main

import (
	"ginSys/common"
	"ginSys/routes"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {

	//初始化db
	db := common.InitDB()
	defer db.Close()

	r := gin.Default()

	//加载路由
	r = routes.CollectRoute(r)

	panic(r.Run())
}
