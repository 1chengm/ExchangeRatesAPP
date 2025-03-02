package router

import (
	"exchangeapp/controllers"
	"exchangeapp/middlewares"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine{
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins: 	[]string{"http://localhost:5173","http://47.109.182.169"},//允许所有源
		AllowMethods: 	[]string{"GET","POST","OPTIONS"},//OPTIONS是预检请求
		AllowHeaders: 	[]string{"Origin","Content-Type","Authorization"},
		ExposeHeaders: 	[]string{"Content-Length"},
		AllowCredentials: true,//允许cookie
		MaxAge: 12 * 3600, //缓存预检请求的时间
	}))

	auth := r.Group("/api/auth")
	{
		auth.POST("/login", controllers.Login)
		auth.POST("/register", controllers.Register)
	}
	api := r.Group("/api")
	api.GET("/exchangeRates", controllers.GetExchangeRate)
	api.Use(middlewares.AuthMiddleware())
	{
		api.POST("/exchangeRates", controllers.CreateExchangeRate)
		api.POST("/articles", controllers.CreateArticle)
		api.GET("/articles", controllers.GetArticles)
		api.GET("/articles/:id", controllers.GetArticleByID)
		
		api.POST("/articles/:id/like",controllers.LikeArticle)
		api.GET("/articles/:id/like",controllers.GetArticleLikes)
	}
	return r;
}
