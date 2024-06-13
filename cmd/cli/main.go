package cli

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
	"github.com/vitoraguiardf/golang-quick-start-jwt-auth/cmd"
	"github.com/vitoraguiardf/golang-quick-start-jwt-auth/cmd/http"
	"github.com/vitoraguiardf/golang-quick-start-jwt-auth/internal/core/services"
	"github.com/vitoraguiardf/golang-quick-start-jwt-auth/internal/handlers/authhandler"
	"github.com/vitoraguiardf/golang-quick-start-jwt-auth/internal/repositories/authrepository"
)

var rootCommand = &cobra.Command{
	Use:   "app",
	Short: "Command Line Interface",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Wellcome to the jungle!")
	},
}

func Execute() {
	cmd.StartupEnv()

	rootCommand.AddCommand(&cobra.Command{
		Use:   "serve",
		Short: "Runs http gin server",
		Long:  ``,
		Run: func(cmd *cobra.Command, args []string) {
			http.Run()
		},
	})

	authRepository := authrepository.NewSqliteRepository()
	authService := services.NewAuthService(authRepository)
	authHandler := authhandler.NewCLIHandler(authService)
	authHandler.RegistyCommands(rootCommand)

	if err := rootCommand.Execute(); err != nil {
		log.Fatal(err)
	}
}
