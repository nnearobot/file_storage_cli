/*
Copyright © 2023 Maksiutova Rimma <rimma.maksiutova@gmail.com>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
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
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
}

func initConfig() {
	viper.AddConfigPath(".")
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")

	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		fmt.Printf("error loading config file %s: %s\n", viper.ConfigFileUsed(), err)
	}
}
