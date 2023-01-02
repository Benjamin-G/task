package cmd

import (
	"fmt"
	"strconv"
	"task/db"

	"github.com/spf13/cobra"
)

var doCmd = &cobra.Command{
	Use:   "do",
	Short: "Marks a task as complete",
	Run: func(cmd *cobra.Command, args []string) {
		var ids []int
		for _, arg := range args {
			id, err := strconv.Atoi(arg)
			if err != nil {
				fmt.Println("Failed to parse the argument:", arg)
			} else if id <= 0 {
				fmt.Println("Failed ID must be a positive number:", arg)
			} else {
				ids = append(ids, id)
			}
		}

		for _, id := range ids {
			task, err := db.GetTask(id)
			if err != nil {
				fmt.Println("Invalid task number:", id)
				continue
			}

			err = db.DeleteTask(task.Key)

			if err != nil {
				fmt.Printf("Failed to mark \"%d\" as completed. Error: %s\n", id, err)
			} else {
				fmt.Printf("Marked \"%d\" as completed.\n", id)
			}
		}
	},
}

func init() {
	RootCmd.AddCommand(doCmd)
}
