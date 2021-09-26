package middleware

import (
	"gintest1/common"
	"gintest1/model"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc{
	return func (c *gin.Context){
		//得到授权
		tokenString := c.GetHeader("Authorization")

		if tokenString == "" || !strings.HasPrefix(tokenString,"Bearer "){
			c.JSON(http.StatusUnauthorized,gin.H{
				"code":401,
				"msg":"权限不足",
			})
			c.Abort()
			return 
		}
		tokenString =tokenString[7:]
		token,claims,err := common.ParseToken(tokenString)
		if err !=nil|| !token.Valid{
			c.JSON(http.StatusUnauthorized,gin.H{
				"code":401,
				"msg":"权限不足",
			})
			c.Abort()
			return
		}

		//验证通过后获取claims 中的userId
		userId := claims.UserId
		DB := common.GetDB()
		var user model.User
		DB.First(&user,userId)

		//用户
		if user.ID ==0{
			c.JSON(http.StatusUnauthorized,gin.H{
				"code":401,
				"msg":"权限不足",
			})
			c.Abort()
			return
		}

		//用户存在 将user信息写入上下文
		c.Set("user",user)
		c.Next()
	}
}