package main

import (
	"ginSys/common"
	"ginSys/routes"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {

	db := common.InitDB()
	defer db.Close()

	r := gin.Default()

	r = routes.CollectRoute(r)

	panic(r.Run())
}
