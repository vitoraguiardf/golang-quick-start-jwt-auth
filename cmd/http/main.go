package main

import (
	"github.com/gin-gonic/gin"
	"github.com/vitoraguiardf/golang-quick-start-jwt-auth/internal/core/services"
	usershandler "github.com/vitoraguiardf/golang-quick-start-jwt-auth/internal/handlers/usersHandler"
	usersrepository "github.com/vitoraguiardf/golang-quick-start-jwt-auth/internal/repositories/usersRepository"
)

func main() {
	router := gin.Default()

	userRepository := usersrepository.NewRepository()
	userService := services.NewUserService(userRepository)
	userHandler := usershandler.NewHttpHandler(userService)
	userHandler.RegistryRoutes(router)

	router.Run(":8080")
}
