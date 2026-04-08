/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"my-cli/internal/todo"
	"strconv"

	"github.com/spf13/cobra"
)

// markCompleteCmd represents the markComplete command
var markCompleteCmd = &cobra.Command{
	Use:   "markComplete [task_id]",
	Short: "Marks task [task_id] as complete",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		id, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Println("invalid integer")
			return
		}
		err = todo.MarkComplete(id)
		if err != nil {
			fmt.Println("error:", err)
		}
		fmt.Println("task completed")
	},
}

func init() {
	rootCmd.AddCommand(markCompleteCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// markCompleteCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// markCompleteCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
