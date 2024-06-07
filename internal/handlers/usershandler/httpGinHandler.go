package usershandler

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/vitoraguiardf/golang-quick-start-jwt-auth/internal/core/ports"
)

type HTTPGinHandler struct {
	service ports.UserService
}

func NewHttpHandler(userService ports.UserService) *HTTPGinHandler {
	return &HTTPGinHandler{
		service: userService,
	}
}

func (handler *HTTPGinHandler) RegistryRoutes(router *gin.Engine) {
	router.GET("/users", handler.Get) // TODO
	router.GET("/users/:id", handler.Get)
	// router.POST("/games", handler.Create) // TODO
}

func (handler *HTTPGinHandler) Get(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.AbortWithStatusJSON(500, gin.H{"error": err.Error()})
	}
	user, err := handler.service.Get(uint(id))
	if err != nil {
		c.AbortWithStatusJSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, user)
}
