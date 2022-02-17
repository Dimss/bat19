package main

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
	"strings"
)

var rootCmd = &cobra.Command{
	Use:   "bat19",
	Short: "bat19 - the game of the year (2019)",
	Run: func(cmd *cobra.Command, args []string) {
		log.Info("start")
		NewScene().play()
	},
}

func init() {
	cobra.OnInitialize(initConfig)

}

func initConfig() {
	viper.AutomaticEnv()
	viper.SetEnvPrefix("T")
	viper.SetEnvKeyReplacer(strings.NewReplacer("-", "_"))
	log.SetLevel(log.InfoLevel)
	log.SetFormatter(&log.TextFormatter{FullTimestamp: true})
	log.SetOutput(os.Stdout)
}

func main() {

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

}
