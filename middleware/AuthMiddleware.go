package middleware

import (
	"ginSys/common"
	"ginSys/model"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		//获取Authorization header
		tokenString := ctx.GetHeader("Authorization")

		//校验 token 格式
		if tokenString == "" || !strings.HasPrefix(tokenString, "Bearer ") {
			ctx.JSON(http.StatusUnauthorized, gin.H{"code": 402,"msg": "权限不足"})
			ctx.Abort()
			return
		}
		tokenString = tokenString[7:]

		token, claims, err := common.ParseToken(tokenString)
		if err != nil || !token.Valid {
			ctx.JSON(http.StatusUnauthorized,gin.H{"code": 401,"msg": "权限不足"})
			ctx.Abort()
			return
		}

		//验证通过后 获取Claim中的ID
		userId := claims.Id
		DB := common.GetDB()
		var user  model.User
		DB.First(&user,userId)

		if user.ID == 0 {
			ctx.JSON(http.StatusUnauthorized,gin.H{"code": 401,"msg": "权限不足"})
			ctx.Abort()
			return
		}

		//用户存在 将user信息写入上下文
		ctx.Set("user", user)
		ctx.Next()
	}
}
