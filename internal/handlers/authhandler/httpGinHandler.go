package authhandler

import (
	"github.com/gin-gonic/gin"
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

func (handler *HTTPGinHandler) RegistryRoutes(router *gin.Engine) {
	authRouter := router.Group("/auth")
	authRouter.POST("/login", handler.Login)
	authRouter.POST("/refresh", handler.Refresh)
	authRouter.POST("/me", handler.Me)
	authRouter.POST("/logout", handler.Logout)
}

func (handler *HTTPGinHandler) Login(c *gin.Context) {
	// TODO
	c.AbortWithStatusJSON(503, gin.H{"error": "Feature in Development"})
}

func (handler *HTTPGinHandler) Refresh(c *gin.Context) {
	// TODO
	c.AbortWithStatusJSON(503, gin.H{"error": "Feature in Development"})
}

func (handler *HTTPGinHandler) Me(c *gin.Context) {
	// TODO
	c.AbortWithStatusJSON(503, gin.H{"error": "Feature in Development"})
}

func (handler *HTTPGinHandler) Logout(c *gin.Context) {
	// TODO
	c.AbortWithStatusJSON(503, gin.H{"error": "Feature in Development"})
}
