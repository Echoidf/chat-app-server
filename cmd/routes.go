package main

import (
	"chat-app/internal/api/controllers"
	"chat-app/pkg/middleware"
	"github.com/gin-gonic/gin"
)

var r *gin.Engine

func init() {
	r = gin.Default()

	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	setupAuthRoutes(r)
	setupUserRoutes(r)
	setupMessageRoutes(r)
}

func setupAuthRoutes(router *gin.Engine) {
	authRoutes := router.Group("/api/auth")
	authRoutes.POST("/login", controllers.Login)
	authRoutes.POST("/signup", controllers.Signup)
	authRoutes.POST("/logout", controllers.Logout)
}

func setupUserRoutes(router *gin.Engine) {
	userRoutes := router.Group("/api/user")
	userRoutes.Use(middleware.JWTAuth())
	userRoutes.GET("/list", controllers.GetUsers)
}

func setupMessageRoutes(router *gin.Engine) {
	messageRoutes := router.Group("/api/message")
	messageRoutes.Use(middleware.JWTAuth())
	messageRoutes.POST("/send/:id", controllers.SendMessage)
	messageRoutes.GET("/:id", controllers.GetMessages)
}
