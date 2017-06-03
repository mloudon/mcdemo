package cmd

import (
	"github.com/mloudon/mcdemo/cache"
	"github.com/spf13/cobra"
)

var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Get a stored file.",
	Long:  `Get a stored file. Two positions arguments - the original filename and the output file to write to`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 2 {
			panic("Must specify stored filename and output as arguments")
		}

		fc := cache.NewFileCacher(mc)
		log.Info("retrieving %s, saving to file %s...", args[0], args[1])
		err := fc.Retrieve(args[0], args[1])
		if err != nil {
			panic(err)
		} else {
			log.Info("...done")
		}
	},
}

func init() {
	RootCmd.AddCommand(getCmd)

}
