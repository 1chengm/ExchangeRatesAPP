package middlewares

import (
	"net/http"
	"exchangeapp/utils"
	"github.com/gin-gonic/gin"
)
func AuthMiddleware() gin.HandlerFunc {
	return func( ctx *gin.Context){
		token := ctx.GetHeader("Authorization")
		if token == ""{
			ctx.JSON(http.StatusUnauthorized,gin.H{"error":"token is missing"})
			ctx.Abort()
			return
		}
		username,err := utils.ParseJWT(token)

			if err != nil{
				ctx.JSON(http.StatusUnauthorized,gin.H{"error":err.Error()})
				ctx.Abort()
				return
			}
		
		ctx.Set("username",username)
		ctx.Next()
	}
}