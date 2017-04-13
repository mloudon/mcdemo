package cmd

import (
	"github.com/mloudon/mcdemo/cache"
	"github.com/spf13/cobra"
)

var putCmd = &cobra.Command{
	Use:   "put",
	Short: "A brief description of your command",
	Long:  ``,

	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			panic("Must specify filename as argument")
		}

		c := cache.New(mc)
		err := c.PutFile(args[0])
		if err != nil {
			panic(err)
		}
	},
}

func init() {
	RootCmd.AddCommand(putCmd)
}
