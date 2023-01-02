/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/spf13/cobra"
)

var version = "0.0.1"

// rootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:     "task",
	Version: version,
	Short:   "Task is a CLI task manager",
}
