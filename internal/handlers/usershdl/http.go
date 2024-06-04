package usershdl

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/vitoraguiardf/golang-quick-start-jwt-auth/internal/core/ports"
)

type userHttpHandler struct {
	userService ports.UserService
}

func NewHttpHandler(userService ports.UserService) *userHttpHandler {
	return &userHttpHandler{
		userService: userService,
	}
}

func (handler *userHttpHandler) RegistryRoutes(router *gin.Engine) {
	router.GET("/users", handler.Get) // TODO
	router.GET("/users/:id", handler.Get)
	// router.POST("/games", handler.Create) // TODO
}

func (handler *userHttpHandler) Get(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.AbortWithStatusJSON(500, gin.H{"error": err.Error()})
	}
	user, err := handler.userService.Get(uint(id))
	if err != nil {
		c.AbortWithStatusJSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, user)
}
