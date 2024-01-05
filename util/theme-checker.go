package util

import (
	"os"

	"github.com/spf13/viper"
)

func CheckTheme(themeIf interface{}) map[string]bool {
	theme := themeIf.(map[string]interface{})
	gtk := checkGTK(theme[string("gtk")].(string))
	kitty := checkKitty(theme[string("kitty")].(string))
	nvim := checkNvim(theme[string("nvim")].(string))

	checkMap := make(map[string]bool)
	checkMap["gtk"] = gtk
	checkMap["kitty"] = kitty
	checkMap["nvim"] = nvim

	return checkMap
}

func checkGTK(theme string) bool {
	if _, err := os.Stat(os.Getenv("HOME") + "/.themes/" + theme); os.IsNotExist(err) {
		return false
	}
	return true
}

func checkKitty(theme string) bool {
	if _, err := os.Stat(os.Getenv("HOME") + "/.config/kitty/" + theme + ".conf"); os.IsNotExist(err) {
		return false
	}
	return true
}

func checkNvim(theme string) bool {
	if _, err := os.Stat(viper.GetString("nvim-plugins-dir") + theme + ".lua"); os.IsNotExist(err) {
		return false
	}
	return true
}
