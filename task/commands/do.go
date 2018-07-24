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
				fmt.Println("Failed to parse the argument:", arg)
			} else {
				ids = append(ids, id)
			}
		}
		tasks, err := db.GetAllTasks()
		if err != nil {
			fmt.Println("Something went wrong:", err)
			return
		}
		for _, id := range ids {
			if id <= 0 || id > len(tasks) {
				fmt.Println("Invalid task number:", id)
				continue
			}
			task := tasks[id-1]
			err := db.DeleteTasks(task.Key)
			fmt.Println(id)
			if err != nil {
				fmt.Printf("Failed to mark \"%d\" as completed. Error: %s\n", id, err)
			} else {
				fmt.Printf("Marked \"%d\" as completed.\n", id)
			}
		}
	},
}

func init() {
	RootCmd.AddCommand(doCommand)
}
