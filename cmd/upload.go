/*
Copyright © 2023 Maksiutova Rimma <rimma.maksiutova@gmail.com>
*/
package cmd

import (
	"fmt"

	"github.com/nnearobot/fstorage/pkg/filestorage"
	"github.com/spf13/cobra"
)

// uploadCmd represents the upload command
var uploadCmd = &cobra.Command{
	Use:   "upload",
	Short: "Upload a file to the filestorage server",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			fmt.Println("Please specify the file to upload.")
			return
		}

		var filePath string = args[0]

		filestorage.UploadFile(filePath)
	},
}

func init() {
	rootCmd.AddCommand(uploadCmd)
}
