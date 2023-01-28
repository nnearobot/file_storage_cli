/*
Copyright © 2023 Maksiutova Rimma <rimma.maksiutova@gmail.com>
*/
package cmd

import (
	"fmt"

	"github.com/nnearobot/fstorage/pkg/filestorage"
	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all the files located on the filestorage server",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fileList, err := filestorage.GetFileList()

		if err != nil {
			fmt.Printf("Error occured while listing a files: %s\n", err)
			return
		}

		fmt.Print("******* Files on the server: *******\n\n")
		for _, fileName := range fileList {
			fmt.Println(fileName)
		}
		fmt.Print("\n")
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
