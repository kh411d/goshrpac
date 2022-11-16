package cmd

import (
	"log"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "goshrpac",
	Short: "goshrpac",
	Long:  `goshrpac`,
}

func init() {
	cobra.OnInitialize(viper.AutomaticEnv)
}

// New Initialize registered cli commands
func New() *cobra.Command {

	rootCmd.Run = func(cmd *cobra.Command, args []string) {
		log.Println("hello world")
	}

	return rootCmd
}
