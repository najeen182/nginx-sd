package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/najeen182/nginx-sd/config"
	"github.com/spf13/cobra"
)

var cfgFile string
var C config.Config

var rootCmd = &cobra.Command{
	Use:   "nginx-sd",
	Short: "Nginx service discovery",
}

// Execute : Execute Cli Command
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.nginx-sd.toml)")
}

func initConfig() {
	if &cfgFile == nil {
		os.Exit(1)
	}
	config.RuntimeViper.SetConfigFile(cfgFile)
	config.RuntimeViper.SetConfigType("toml")
	if err := config.RuntimeViper.ReadInConfig(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	err := config.RuntimeViper.Unmarshal(&C)
	if err != nil {
		log.Println(err)

	}

}
