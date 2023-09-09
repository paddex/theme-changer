/*
Copyright © 2023 Patrik Dufeu <paddex@paddex.net>
*/
package cmd

import (
	"log"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	version = "0.1.0"
	commit  = "none"
	date    = "unknown"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:     "theme",
	Version: version,
	Short:   "This tool switches themes including GTK2/3/4, kitty and nvim",
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func addMyCommands() {
	rootCmd.AddCommand(listCmd)
	rootCmd.AddCommand(switchCmd)
}

func init() {
	rootCmd.SetVersionTemplate("Version: " + version + "\nCommit: " + commit + "\nDatum: " + date + "\n")

	viper.SetConfigName("theme-changer")
	viper.SetConfigType("json")
	viper.AddConfigPath("/etc/theme-changer/")
	viper.AddConfigPath("$HOME/.config/theme-changer/")
	viper.AddConfigPath(".")

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal("Fatal error reading config file: %w", err)
	}

	addMyCommands()
}
