package controller

import (
	"ginSys/common"
	"ginSys/model"
	"ginSys/util"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
)

//注册用户
func Register(ctx *gin.Context) {
	DB := common.GetDB()
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
		name = util.RandomString(10)
	}
	log.Println(name, telephone, password)
	//判断手机号是否存在
	if isTelephoneExists(DB, telephone) {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "手机号已存在"})
		return
	}
	//创建用户
	hasedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError,gin.H{"code": 500,"msg": "加密错误"})
	}
	newUser := model.User{
		Name:      name,
		Telephone: telephone,
		Password:  string(hasedPassword),
	}
	DB.Create(&newUser)
	//数据返回
	ctx.JSON(200, gin.H{
		"code": 200,
		"message": "注册成功",
	})
}
//用户登录
func Login(ctx *gin.Context)  {
	DB := common.GetDB()
	//获取参数
	password := ctx.PostForm("password")
	telephone := ctx.PostForm("telephone")
	//参数验证
	if len(telephone) != 11 {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "手机号必须11位"})
		return
	}
	if len(password) < 6 {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "密码必须大于6位"})
		return
	}
	//手机号是否存在
	var user model.User
	DB.Where("telephone = ?", telephone).First(&user)
	if user.ID == 0 {
		ctx.JSON(http.StatusUnprocessableEntity,gin.H{"code": 422, "msg": "用户不存在！"})
		return
	}

	//密码是否正确
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)) ; err != nil {
		ctx.JSON(http.StatusInternalServerError,gin.H{"code": 500,"msg": "密码不正确!"})
	}
	//发放token
	token, err := common.ReleaseToken(user)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError,gin.H{"code": 500, "msg": "系统异常"})
		log.Printf("token generate error : %v", err)
		return
	}

	ctx.JSON(200,gin.H{
		"code": 200,
		"data": gin.H{"token": token},
		"msg": "登录成功",
	})
}

//获取用户信息
func Info(ctx *gin.Context)  {
	user, _ := ctx.Get("user")

	ctx.JSON(http.StatusOK,gin.H{"code": 200,"data": gin.H{"user": user}})
}

func isTelephoneExists(db *gorm.DB, telephone string) bool {
	var user model.User
	db.Where("telephone = ?", telephone).First(&user)
	if user.ID != 0 {
		return true
	}

	return false
}
