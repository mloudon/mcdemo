package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var putCmd = &cobra.Command{
	Use:   "put",
	Short: "A brief description of your command",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("put called")
	},
}

func init() {
	RootCmd.AddCommand(putCmd)
}
