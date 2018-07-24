package commands

import (
	"fmt"
	"os"
	"task/db"

	"github.com/spf13/cobra"
)

// listCommand represents the list command used to list all the tasks
// Run() function defined inside lists all the tasks stored in Bolt DB
var listCommand = &cobra.Command{
	Use:     "list",
	Short:   "Lists all your tasks ",
	Example: "task list",
	Run: func(command *cobra.Command, args []string) {
		tasks, err := db.GetAllTasks()
		if err != nil {
			fmt.Println("Error displaying tasks:", err.Error())
			os.Exit(1)
		}
		if len(tasks) == 0 {
			fmt.Printf("No tasks to finish\n")
			return
		}
		fmt.Println("List of Tasks:")
		for i := 0; i < len(tasks); i++ {
			fmt.Printf("%d.%s\n", i+1, tasks[i].Value)
		}
	},
}

func init() {
	RootCmd.AddCommand(listCommand)
}
