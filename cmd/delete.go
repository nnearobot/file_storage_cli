/*
Copyright © 2023 Maksiutova Rimma <rimma.maksiutova@gmail.com>
*/
package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/nnearobot/fstorage/pkg/filestorage"
	"github.com/spf13/cobra"
)

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a file by providing a file name",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			fmt.Println("Please specify exactly one file name.")
			return
		}

		var fileName string = args[0]

		if confirm(fmt.Sprintf("Do you really want to delete the file %s?", fileName), 3) {
			filestorage.DeleteFile(args[0])
			return
		}

		fmt.Println("Abort")
	},
}

// confirm asks user if they are sure or not to execute a command
func confirm(s string, tries int) bool {
	r := bufio.NewReader(os.Stdin)

	for ; tries > 0; tries-- {
		fmt.Printf("%s [y/n]: ", s)

		res, err := r.ReadString('\n')
		if err != nil {
			return false
		}

		if len(res) < 2 {
			continue
		}

		return strings.ToLower(strings.TrimSpace(res))[0] == 'y'
	}

	return false
}

func init() {
	rootCmd.AddCommand(deleteCmd)
}
