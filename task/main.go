package main

import (
	"path/filepath"
	"secret/cli/cmd"

	"task/db"

	"github.com/mitchellh/go-homedir"
)

func main() {
	initApp()
}

func initApp() error {
	h, _ := homedir.Dir()
	dbPath := filepath.Join(h, "tasks.db")
	_, err := db.InitDB(dbPath)
	cmd.RootCmd.Execute()
	return err
}
