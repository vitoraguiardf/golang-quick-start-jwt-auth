package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/vitoraguiardf/golang-quick-start-jwt-auth/internal/core/services"
	"github.com/vitoraguiardf/golang-quick-start-jwt-auth/internal/handlers/usershdl"
	"github.com/vitoraguiardf/golang-quick-start-jwt-auth/internal/repositories/usersrepo"
)

func main() {

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file: %s", err)
	}

	router := gin.Default()

	userRepository := usersrepo.NewRepository()
	userService := services.NewUserService(userRepository)
	userHandler := usershdl.NewHttpHandler(userService)
	userHandler.RegistryRoutes(router)

	router.Run(":8080")
}
