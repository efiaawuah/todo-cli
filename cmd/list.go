/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"my-cli/internal/todo"

	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "list tasks in todo list",
	Long:  `lists tasks in todo list. Flag -complete for all done tasks and -inconplete for all incomplete tasks`,
	Run: func(cmd *cobra.Command, args []string) {

		complete, _ := cmd.Flags().GetBool("complete")
		incomplete, _ := cmd.Flags().GetBool("incomplete")

		var filter *bool

		if complete {
			t := true
			filter = &t
		} else if incomplete {
			f := false
			filter = &f
		} else {
			filter = nil // no filter
		}

		err, tasks := todo.ListTasks(filter)
		if err != nil {
			fmt.Println("error:", err)
		}

		for _, task := range tasks {
			fmt.Printf("%d: %s (done=%v)\n", task.ID, task.DESCRIPTION, task.DONE)
		}

	},
}

func init() {
	rootCmd.AddCommand(listCmd)

	listCmd.Flags().BoolP("complete", "c", false, "show only completed tasks")
	listCmd.Flags().BoolP("incomplete", "i", false, "show only incomplete tasks")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
