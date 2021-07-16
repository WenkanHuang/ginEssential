package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	"xietong.me/ginessential/common"
	"xietong.me/ginessential/model"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		//获取authorization header
		tokenString := c.GetHeader("Authorization")
		prefex := "Todo"
		//validate token
		if tokenString == "" || !strings.HasPrefix(tokenString, prefex) {
			c.JSON(http.StatusUnauthorized, gin.H{"code": 401, "msg": "权限不足"})
			c.Abort()
			return
		}

		tokenString = tokenString[len(prefex)+1:]
		token, claims, err := common.ParseToken(tokenString)
		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"code": 401, "msg": "权限不足"})
			c.Abort()
			return
		}

		//验证通过后获取Claim中的userId
		userId := claims.UserId
		DB := common.GetDB()
		var user model.User
		DB.First(&user, userId)

		//用户不存在
		if user.UserId == 0 {
			c.JSON(http.StatusUnauthorized, gin.H{"code": 401, "msg": "权限不足"})
			c.Abort()
			return
		}
		//用户存在，将user信息写入上下文
		c.Set("user", user)
		c.Next()
	}
}
