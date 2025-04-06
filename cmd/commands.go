package cmd

import (
	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new to-do item",
	Long:  "Add a new to-do item to the list",
	Run:   addRecord,
}

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all uncompleted to-do items",
	Long:  "List all uncompleted to-do items",
	Run:   listRecords,
}

func init() {
	listCmd.Flags().BoolP("all", "a", false, "List all to-do items")

	rootCmd.AddCommand(addCmd)
	rootCmd.AddCommand(listCmd)
}
