package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// archivedCmd represents the archived command
var archivedCmd = &cobra.Command{
	Use:   "archived",
	Short: "archived",
	Long:  `archived`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("archived called")
	},
}

func init() {
	downloadCmd.AddCommand(archivedCmd)
}
