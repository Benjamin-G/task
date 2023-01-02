/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"cmd/cmd"
	"cmd/db"
	"fmt"
	"os"
	"path/filepath"

	"github.com/mitchellh/go-homedir"
)

func main() {
	home, _ := homedir.Dir()
	dbPath := filepath.Join(home, "tasks.db")

	fmt.Println("Starting", dbPath)
	must(db.Init(dbPath))
	must(cmd.RootCmd.Execute())
}

func must(err error) {
	if err != nil {
		fmt.Println("Error: ", err)
		fmt.Println(err.Error())
		os.Exit(1)
	}
}
