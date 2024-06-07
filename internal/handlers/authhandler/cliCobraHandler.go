package authhandler

import (
	"github.com/vitoraguiardf/golang-quick-start-jwt-auth/internal/core/ports"
)

type CLICobraHandler struct {
	authService ports.AuthService
}

func NewCLIHandler(authService ports.AuthService) *CLICobraHandler {
	return &CLICobraHandler{
		authService: authService,
	}
}

func (handler *CLICobraHandler) RegistryCommand() {
}

func (handler *CLICobraHandler) Login() {
	// TODO
}

func (handler *CLICobraHandler) Refresh() {
	// TODO
}

func (handler *CLICobraHandler) Me() {
	// TODO
}

func (handler *CLICobraHandler) Logout() {
	// TODO
}
