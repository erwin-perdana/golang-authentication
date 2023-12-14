package main

import (
	"github.com/gin-gonic/gin"
	"golang-authentication/initializers"
	"golang-authentication/controller"
	"golang-authentication/middleware"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDb()
	initializers.SyncDatabase()
}

func main()  {
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	api := r.Group("/api")
	{
		api.POST("/transactions", middleware.RequireAuth, controller.Buy)
		api.GET("/transactions/:id", middleware.RequireAuth, controller.GetTransactions)
		api.POST("/payments", middleware.RequireAuth, controller.Pay)

		auth := api.Group("/auth")
		{
			auth.POST("/signup", controller.Signup)
			auth.POST("/login", controller.Login)
			auth.POST("/logout", middleware.RequireAuth, controller.Logout)
		}
	}

	r.Run()
}