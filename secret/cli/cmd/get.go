package cmd

import (
	"fmt"

	secret "secret/vault"

	"github.com/spf13/cobra"
)

//handler for get command
var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Gets a secret in your secret storage",
	Run: func(cmd *cobra.Command, args []string) {
		v := secret.NewVault(encodingKey, secretsPath())
		key := args[0]
		value, err := v.GetValue(key)
		if err != nil {
			fmt.Println("no value set")
			return
		}
		fmt.Printf("%s = %s\n", key, value)
	},
}

func init() {
	RootCmd.AddCommand(getCmd)
}
