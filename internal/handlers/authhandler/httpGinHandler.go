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
		c.JSON(http.StatusOK, gin.H{"token": token})
		return
	}
}

func (h *HTTPGinHandler) Refresh(c *gin.Context) {
	// TODO
	c.AbortWithStatusJSON(http.StatusServiceUnavailable, gin.H{"error": "Feature in Development"})
}

func (h *HTTPGinHandler) Me(c *gin.Context) {
	// TODO
	c.AbortWithStatusJSON(http.StatusServiceUnavailable, gin.H{"error": "Feature in Development"})
}

func (h *HTTPGinHandler) Logout(c *gin.Context) {
	// TODO
	c.AbortWithStatusJSON(http.StatusServiceUnavailable, gin.H{"error": "Feature in Development"})
}
