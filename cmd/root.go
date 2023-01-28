/*
Copyright © 2023 Maksiutova Rimma <rimma.maksiutova@gmail.com>
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var (
	version string = "1.0"

	rootCmd = &cobra.Command{
		Use:     "fstorage",
		Short:   "Manage files on a remote filestorage",
		Long:    `This application allows to list, add, and delete files on a remote filestorage`,
		Version: version,
	}
)

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	rootCmd.CompletionOptions.DisableDefaultCmd = true
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
}
