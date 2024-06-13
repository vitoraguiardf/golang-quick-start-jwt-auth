package authhandler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vitoraguiardf/golang-quick-start-jwt-auth/internal/core/domain"
	"github.com/vitoraguiardf/golang-quick-start-jwt-auth/internal/core/ports"
)

type HTTPGinHandler struct {
	authService ports.AuthService
}

func NewHttpHandler(authService ports.AuthService) *HTTPGinHandler {
	return &HTTPGinHandler{
		authService: authService,
	}
}

func (h *HTTPGinHandler) RegistryRoutes(router *gin.Engine) {
	authRouter := router.Group("/auth")
	authRouter.POST("/login", h.Login)
	authRouter.POST("/refresh", h.Refresh)
	authRouter.POST("/me", h.Me)
	authRouter.POST("/logout", h.Logout)
}

func (h *HTTPGinHandler) Login(c *gin.Context) {
	var credentials domain.Credentials
	if err := c.ShouldBind(&credentials); err != nil {
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}
	if token, err := h.authService.Login(credentials); err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	} else {
		c.JSON(http.StatusOK, token)
		return
	}
}

func (h *HTTPGinHandler) Refresh(c *gin.Context) {
	authorization := c.GetHeader("Authorization")
	if token, err := h.authService.Refresh(authorization); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	} else {
		c.JSON(http.StatusOK, token)
		return
	}
}

func (h *HTTPGinHandler) Me(c *gin.Context) {
	authorization := c.GetHeader("Authorization")
	if user, err := h.authService.Me(authorization); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	} else {
		c.JSON(http.StatusOK, user)
		return
	}
}

func (h *HTTPGinHandler) Logout(c *gin.Context) {
	authorization := c.GetHeader("Authorization")
	if message, err := h.authService.Logout(authorization); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{"message": message})
		return
	}
}
