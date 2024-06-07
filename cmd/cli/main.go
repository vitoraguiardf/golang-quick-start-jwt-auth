package cli

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
	"github.com/vitoraguiardf/golang-quick-start-jwt-auth/cmd"
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
	if err := rootCommand.Execute(); err != nil {
		log.Fatal(err)
	}
}
