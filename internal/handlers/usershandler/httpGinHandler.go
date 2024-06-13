package usershandler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vitoraguiardf/golang-quick-start-jwt-auth/internal/core/domain"
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
	router.GET("/users", handler.FindAll)
	router.POST("/users", handler.Create)
	router.GET("/users/:id", handler.FindById)
	router.PUT("/users/:id", handler.Replace)
	router.PATCH("/users/:id", handler.Update)
	router.DELETE("/users/:id", handler.Delete)
}

func (handler *HTTPGinHandler) Create(c *gin.Context) {
	var model domain.User
	if err := c.ShouldBind(&model); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}
	if err := handler.service.Create(&model); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"created": fmt.Sprintf("/users/%v", int(model.ID))})
}

func (handler *HTTPGinHandler) FindAll(c *gin.Context) {
	if items, err := handler.service.FindAll(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{"data": &items})
	}
}

func (handler *HTTPGinHandler) FindById(c *gin.Context) {
	if item, err := handler.service.FindById(c.Param("id")); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	} else {
		c.JSON(http.StatusOK, item)
		return
	}
}

func (handler *HTTPGinHandler) Replace(c *gin.Context) {
	var model domain.User
	if err := c.ShouldBind(&model); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}
	if err := handler.service.Replace(c.Param("id"), &model); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	} else {
		c.JSON(http.StatusOK, model)
		return
	}
}

func (handler *HTTPGinHandler) Update(c *gin.Context) {
	var model domain.User
	if err := c.ShouldBind(&model); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}
	if err := handler.service.Update(c.Param("id"), &model); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	} else {
		c.JSON(http.StatusOK, model)
		return
	}
}

func (handler *HTTPGinHandler) Delete(c *gin.Context) {
	if err := handler.service.Delete(c.Param("id")); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	} else {
		c.JSON(http.StatusAccepted, nil)
		return
	}
}
