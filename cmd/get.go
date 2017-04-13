package cmd

import (
	"github.com/Hearst-DD/cache"
	"github.com/spf13/cobra"
)

var getCmd = &cobra.Command{
	Use:   "get",
	Short: "A brief description of your command",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			panic("Must specify filename as argument")
		}

		c := cache.New(mc)
		log.Info("writing %s")
		c.GetFile
	},
}

func init() {
	RootCmd.AddCommand(getCmd)

}
