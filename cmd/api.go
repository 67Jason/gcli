package cmd

import (
	"github.com/spf13/cobra"
	"log"
)

// apiCmd represents the api command
var apiCmd = &cobra.Command{
	Use:   "api",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		initApi(args)
	},
}

func init() {
	rootCmd.AddCommand(apiCmd)
}

func initApi(args []string) {
	if len(args) > 0 {
		log.Println("arg is :", args[0])
	} else {
		log.Println("no args")
	}
}
