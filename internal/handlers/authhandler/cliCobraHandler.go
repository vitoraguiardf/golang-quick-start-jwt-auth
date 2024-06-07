package authhandler

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"syscall"

	"github.com/spf13/cobra"
	"github.com/vitoraguiardf/golang-quick-start-jwt-auth/internal/core/ports"
	"golang.org/x/term"
)

type CLICobraHandler struct {
	authService ports.AuthService
}

func NewCLIHandler(authService ports.AuthService) *CLICobraHandler {
	return &CLICobraHandler{
		authService: authService,
	}
}

var rootCommand = &cobra.Command{
	Use:   "auth",
	Short: "Commands for Auth",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
	},
}

func (handler *CLICobraHandler) RegistyCommands(parentCommand *cobra.Command) {
	parentCommand.AddCommand(rootCommand)

	rootCommand.AddCommand(handler.loginCommand())
	rootCommand.AddCommand(handler.meCommand())
	rootCommand.AddCommand(handler.refreshCommand())
	rootCommand.AddCommand(handler.logoutCommand())
}

func (handler *CLICobraHandler) loginCommand() *cobra.Command {
	var username string
	var loginCommand = &cobra.Command{
		Use:   "login [username]",
		Short: "Login Command",
		Long:  ``,
		Args:  cobra.ExactArgs(0),
		Run: func(cmd *cobra.Command, args []string) {
			reader := bufio.NewReader(os.Stdin)
			if len(username) <= 0 {
				fmt.Println("username:")
				if value, err := reader.ReadString('\n'); err == nil {
					username = value
				} else {
					log.Fatal(err)
				}
			}
			fmt.Println("password:")
			if bytePassword, err := term.ReadPassword(int(syscall.Stdin)); err == nil {
				password := strings.TrimSpace(string(bytePassword))
				fmt.Printf("Running Login to %v!\n", username)
				if token, err := handler.authService.Login(username, password); err != nil {
					fmt.Printf("Login fails with error: %v\n", err)
				} else {
					fmt.Printf("Login successful!\nToken: %v\n", token)
				}
			} else {
				log.Fatal(err)
			}
		},
	}
	loginCommand.Flags().StringVarP(&username, "username", "u", "", "username for login")
	return loginCommand
}

func (handler *CLICobraHandler) refreshCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "refresh",
		Short: "Refresh Command",
		Long:  ``,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Refresh!")
		},
	}
}

func (handler *CLICobraHandler) meCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "me",
		Short: "Me Command",
		Long:  ``,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Me!")
		},
	}
}

func (handler *CLICobraHandler) logoutCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "logout",
		Short: "Logout Command",
		Long:  ``,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Logout!")
		},
	}
}
