package main

import (
	"github.com/gin-gonic/gin"
	"go_free/Controller"
	"go_free/Middleware"
)

func main() {
	//gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	user := r.Group("/api/user")
	{
		user.POST("/register", Controller.Register)
		user.PUT("/login", Controller.Login)
		user.GET("/info", Middleware.AuthMiddleware(), Controller.Info)
		user.DELETE("/logout",Middleware.AuthMiddleware(),Controller.Logout)
	}
	book := r.Group("/api/book")
	{
		book.GET("/", Controller.Allbooks)
		book.POST("/", Controller.Addbook)
		book.PUT("/", Controller.Changebook)
		book.DELETE("/", Controller.Deletebook)
	}
	lend := r.Group("/api/lend")
	{
		lend.POST("/", Controller.Lend)
		lend.DELETE("/", Controller.Back)
	}
	panic(r.Run(":8080"))
}
