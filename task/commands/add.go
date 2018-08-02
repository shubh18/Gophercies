package commands

import (
	"fmt"
	"strings"
	"task/db"

	"github.com/spf13/cobra"
)

//addCommand adds tasks to the list and stores it in Bolt DB
var addCommand = &cobra.Command{
	Use:     "add",
	Short:   "Add is to generate list of tasks",
	Example: "task add Read Golang Tutorial",
	Run: func(command *cobra.Command, args []string) {
		task := strings.Join(args, " ")
		err := db.CreateTask(task)
		msg := "Error in Add task"
		if err == nil {
			msg = fmt.Sprintf("Added %s task to your list\n", task)
		}
		fmt.Printf(msg)
	},
}

func init() {
	RootCmd.AddCommand(addCommand)
}
