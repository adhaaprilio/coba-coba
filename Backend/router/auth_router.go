package router

import (
	"backend/config"
	"backend/handler"
	"backend/middleware"
	"backend/repository"
	"backend/service"

	"github.com/gin-gonic/gin"
)

func AuthRouter(api *gin.RouterGroup) {
	authRepository := repository.NewAuthRepository(config.ConnectDatabase())
	authService := service.NewAuthService(authRepository)
	authHandler := handler.NewAuthHandler(authService)

	api.Use(middleware.Auth())
	api.POST("/user/register", authHandler.Register)
	api.POST("/user/login", authHandler.Login)

}
