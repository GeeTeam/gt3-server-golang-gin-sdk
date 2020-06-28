package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"gt3-server-golang-gin-sdk/controllers"
	"gt3-server-golang-gin-sdk/controllers/sdk"
)

func main() {
	r := gin.Default()
	// 设置静态资源
	r.Static("/static", "./static")
	r.StaticFile("/", "./static/index.html")
	r.StaticFile("/favicon.ico", "./static/favicon.ico")
	// 设置session中间件
	store := cookie.NewStore([]byte(sdk.VERSION))
	r.Use(sessions.Sessions("mysession", store))
	// 设置web路由
	r.GET("/register", controllers.FirstRegister)
	r.POST("/validate", controllers.SecondValidate)
	r.Run(":8000")
}
