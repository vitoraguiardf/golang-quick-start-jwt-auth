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
	var id uint
	if param_id := c.Param("id"); len(param_id) > 0 {
		if key, err := strconv.Atoi(param_id); err == nil {
			id = uint(key)
		} else {
			c.AbortWithStatusJSON(500, gin.H{"error": err.Error()})
			return
		}
	}

	if id != 0 {
		if user, err := handler.service.Get(uint(id)); err != nil {
			c.AbortWithStatusJSON(500, gin.H{"error": err.Error()})
			return
		} else {
			c.JSON(200, user)
			return
		}
	}
	c.JSON(200, gin.H{"message": "listing all users"})
}
