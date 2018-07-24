package main

import (
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
	err := db.InitDB(DbPath)
	if err != nil {
		panic(err)
	}
	return err
}
