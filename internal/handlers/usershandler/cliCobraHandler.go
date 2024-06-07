package usershandler

import (
	"github.com/vitoraguiardf/golang-quick-start-jwt-auth/internal/core/ports"
)

type CLICobraHandler struct {
	service ports.UserService
}

func NewCLIHandler(userService ports.UserService) *CLICobraHandler {
	return &CLICobraHandler{
		service: userService,
	}
}

func (handler *CLICobraHandler) RegistryCommand() {
	// TODO
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
