package main

import (
	"fmt"
	"path/filepath"
	"task/commands"
	"task/db"

	homedir "github.com/mitchellh/go-homedir"
)

func main() {
	initApplication()
	commands.RootCmd.Execute()

}

func initApplication() error {
	home, _ := homedir.Dir()
	DbPath := filepath.Join(home, "cmd.db")
	_, err := db.InitDB(DbPath)
	if err != nil {
		fmt.Println(err)
	}
	return err
}
