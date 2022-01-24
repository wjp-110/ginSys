package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"log"
	"math/rand"
	"net/http"
	"time"
)

type User struct {
	gorm.Model
	Name      string `gorm:"type:varchar(20);not null"`
	Telephone string `gorm:"type:varchar(110);not null;unique"`
	Password  string `gorm:type:varchar(225);not null`
}

func main() {

	db := initDB()
	defer db.Close()

	r := gin.Default()
	r.POST("/api/author/register", func(ctx *gin.Context) {
		//获取参数
		name := ctx.PostForm("name")
		password := ctx.PostForm("password")
		telephone := ctx.PostForm("telephone")

		//数据验证
		if len(telephone) != 11 {
			ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "手机号必须11位"})
			return
		}
		if len(password) < 6 {
			ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "密码必须大于6位"})
			return
		}

		//如果名称没有传，给一个10位的随机字符串
		if len(name) == 0 {
			name = RandomString(10)
		}

		log.Println(name, telephone, password)

		//判断手机号是否存在
		if isTelephoneExists(db, telephone) {
			ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "手机号已存在"})
			return
		}
		//创建用户
		newUser := User{
			Name:      name,
			Telephone: telephone,
			Password:  password,
		}

		db.Create(&newUser)

		//数据返回
		ctx.JSON(200, gin.H{
			"message": "注册成功",
		})
	})
	panic(r.Run())
}

func isTelephoneExists(db *gorm.DB, telephone string) bool {
	var user User
	db.Where("telephone = ?", telephone).First(&user)
	if user.ID != 0 {
		return true
	}

	return false
}

func RandomString(n int) string {
	var letters = []byte("asvdgahahsbdagaskdashasdknjsal")
	result := make([]byte, n)

	rand.Seed(time.Now().Unix())
	for i := range result {
		result[i] = letters[rand.Intn(len(letters))]
	}

	return string(result)
}

func initDB() *gorm.DB {
	db, err := gorm.Open("mysql", "root:123456@/gin_sys?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("failed to connect database, err---->" + err.Error())
	}
	db.AutoMigrate(&User{})
	return db
}
