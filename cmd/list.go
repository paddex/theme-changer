package cmd

import (
	"fmt"

	"paddex.net/theme-changer/util"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Lists all available themes",
	Run: func(cmd *cobra.Command, args []string) {
		themes := viper.GetViper().GetStringMap("themes")

		for key, values := range themes {
			checkMap := util.CheckTheme(values)
			if !checkMap["gtk"] || !checkMap["kitty"] || !checkMap["nvim"] {
				fmt.Println(util.Red + key + util.Reset)
				continue
			}
			fmt.Println(util.Green + key + util.Reset)
		}
	},
}
