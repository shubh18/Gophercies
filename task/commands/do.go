package commands

import (
	"fmt"
	"strconv"
	"task/db"

	"github.com/spf13/cobra"
)

// doCommand represents the do command
// doCommand used to mark task as complete and it deletes
var doCommand = &cobra.Command{
	Use:     "do",
	Short:   "Marks a task as complete",
	Example: "task do task_number \n eg. task do 1",
	Run: func(command *cobra.Command, args []string) {
		var ids []int
		for _, arg := range args {
			id, err := strconv.Atoi(arg)
			if err != nil {
				return
			}
			ids = append(ids, id)
		}
		tasks, err := db.GetAllTasks()
		if err != nil {
			fmt.Println("error occured")
		}
		for _, id := range ids {
			if id <= 0 || id > len(tasks) {
				fmt.Println("Invalid task number:", id)
				continue
			}
			task := tasks[id-1]
			err := db.DeleteTasks(task.Key)
			msg := "Failed to delete"
			if err == nil {
				msg = fmt.Sprintf("Marked \"%d\" as completed.\n", id)
			}
			fmt.Printf(msg)

		}
	},
}

func init() {
	RootCmd.AddCommand(doCommand)
}
