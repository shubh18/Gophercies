package commands

import (
	"github.com/spf13/cobra"
)

// RootCmd - definition of root command
// We
var RootCmd = &cobra.Command{
	Use:   "task",
	Short: "Task is CLI manager",
}
