package main

import (
	"github.com/gin-gonic/gin"
	"github.com/vitoraguiardf/golang-quick-start-jwt-auth/cmd"
	"github.com/vitoraguiardf/golang-quick-start-jwt-auth/internal/core/services"
	"github.com/vitoraguiardf/golang-quick-start-jwt-auth/internal/handlers/authhandler"
	"github.com/vitoraguiardf/golang-quick-start-jwt-auth/internal/handlers/usershandler"
	"github.com/vitoraguiardf/golang-quick-start-jwt-auth/internal/repositories/authrepository"
	"github.com/vitoraguiardf/golang-quick-start-jwt-auth/internal/repositories/usersrepo"
)

func main() {
	cmd.StartupEnv()

	router := gin.Default()

	authRepository := authrepository.NewKvsRepository()
	authService := services.NewAuthService(authRepository)
	authHandler := authhandler.NewHttpHandler(authService)
	authHandler.RegistryRoutes(router)

	userRepository := usersrepo.NewKvsRepository()
	userService := services.NewUserService(userRepository)
	userHandler := usershandler.NewHttpHandler(userService)
	userHandler.RegistryRoutes(router)

	router.Run(":8080")
}
