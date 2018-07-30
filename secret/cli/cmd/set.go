package cmd

import (
	"fmt"

	secret "secret/vault"

	"github.com/spf13/cobra"
)

var setCmd = &cobra.Command{
	Use:   "set",                                  // Name of the command
	Short: "Sets a secret in your secret storage", // Command description in short
	//anonymous function
	Run: func(cmd *cobra.Command, args []string) {
		v := secret.NewVault(encodingKey, secretsPath())
		key, value := args[0], args[1]
		err := v.SetKey(key, value)
		if err != nil {
			return
		}
		fmt.Println("Value set successfully!")
	},
}

func init() {
	RootCmd.AddCommand(setCmd)
}
