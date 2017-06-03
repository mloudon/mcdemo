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

		fc := cache.NewFileCacher(mc)
		log.Info("storing %s...", args[0])
		err := fc.Store(args[0])
		if err != nil {
			panic(err)
		} else {
			log.Info("...done")
		}
	},
}

func init() {
	RootCmd.AddCommand(putCmd)
}
