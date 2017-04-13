package cmd

import (
	"fmt"
	"os"

	"github.com/bradfitz/gomemcache/memcache"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	l5g "github.com/neocortical/log5go"
)

var log = l5g.Logger(l5g.LogAll)

var (
	cfgFile string
	mc      *memcache.Client
)

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "mcdemo",
	Short: "Store and retrieve large files from memcached",
	Long:  `Store and retrieve large files from memcached`,
}

func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	RootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.mcdemo.yaml)")

	mc = memcache.New(fmt.Sprintf("%s:%d", viper.GetString("mchost"), viper.GetInt("mcport")))
	log.Debug("mc client: %+v", mc)
}

func initConfig() {
	if cfgFile != "" { // enable ability to specify config file via flag
		viper.SetConfigFile(cfgFile)
	}

	viper.SetConfigName(".mcdemo") // name of config file (without extension)
	viper.AddConfigPath("$HOME")   // adding home directory as first search path
	viper.AutomaticEnv()           // read in environment variables that match

	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
