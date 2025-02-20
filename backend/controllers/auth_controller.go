package controllers

import (
    "exchangeapp/models"
    "exchangeapp/utils"
	"exchangeapp/global"
    "net/http"
    "github.com/gin-gonic/gin"

)

func Register(ctx *gin.Context) {
    var user models.User
    
    if err := ctx.ShouldBindJSON(&user); err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
	//将请求的json数据绑定到user结构体上，如果绑定失败返回错误
    hashedPwd, err := utils.HashPassword(user.Password)//加密密码
    if err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    user.Password = hashedPwd
    token, err := utils.GenerateJWT(user.Username)//生成token,jwt
    if err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
	if err := global.Db.AutoMigrate(&user); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}//创建表,如果表已经存在则不会创建
	if err := global.Db.Create(&user).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError,gin.H{"error":err.Error()})
		return
	}//插入数据
    ctx.JSON(http.StatusOK, gin.H{"token": token})
}
func Login(ctx *gin.Context) {
	var input struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var user models.User
	//从数据库中查找用户,如果查不到说明用户不存在
	if err := global.Db.Where("username = ?", input.Username).First(&user).Error; err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "invalid user"})
		return
	}
	if !utils.CheckPasswordHash(input.Password, user.Password){
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "invalid password"})
	}//检查密码是否正确
	//生成token,如果用户名和密码都正确，生成与用户注册时相同的token
	token, err := utils.GenerateJWT(user.Username)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"token": token})
}