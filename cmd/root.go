package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "To-do app",
	Short: "A simple CLI to-do application",
	Long:  "A simple CLI to-do application that allows you to manage your tasks.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Welcome to the To-do app!")
	},
}

func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}
