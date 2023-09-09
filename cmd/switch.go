package cmd

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"

	"paddex.net/theme-changer/util"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var switchCmd = &cobra.Command{
	Use:   "switch <theme>",
	Short: "Switches to the given theme",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		theme := args[0]
		themeMap := viper.GetStringMap("themes." + theme)
		checkMap := util.CheckTheme(themeMap)
		if !checkMap["gtk"] || !checkMap["kitty"] || !checkMap["nvim"] {
			fmt.Println("Theme is not fully installed. Aborting...")
			return
		}

		switchGtk(themeMap[string("gtk")].(string))
		switchShell(themeMap[string("shell")].(string))
		switchNvim(themeMap[string("nvim")].(string))
		switchKitty(themeMap[string("kitty")].(string))
	},
}

func switchGtk(theme string) {
	// Change GTK2/3 Command:
	// gsettings set org.gnome.desktop.interface gtk-theme
	cmd := exec.Command("/usr/bin/gsettings", "set", "org.gnome.desktop.interface", "gtk-theme", theme)

	err := cmd.Run()
	if err != nil {
		log.Fatal("Was not able to execute command to change GTK-Theme: ", err)
	}

	err = os.Remove(filepath.Join(os.Getenv("HOME"), ".config/gtk-4.0"))
	if err != nil {
		log.Fatal("Unable to remove gtk-4.0 symlink: ", err)
	}
	err = os.Symlink(filepath.Join(os.Getenv("HOME"), ".themes", theme, "gtk-4.0"), filepath.Join(os.Getenv("HOME"), ".config/gtk-4.0"))
	if err != nil {
		log.Fatal("Unable to create new gtk-4.0 symlink: ", err)
	}
}

func switchShell(theme string) {
	// Change Shell Command:
	// gsettings set org.gnome.shell.extensions.user-theme name
	cmd := exec.Command("/usr/bin/gsettings", "set", "org.gnome.shell.extensions.user-theme", "name", theme)

	err := cmd.Run()
	if err != nil {
		log.Fatal("Was not able to execute command to change shell: ", err)
	}
}

func switchNvim(theme string) {
	csFileName := viper.GetString("nvim-colorscheme-file")
	csFileContent, err := os.ReadFile(csFileName)
	if err != nil {
		log.Fatal("Could not open nvim-colorscheme-file")
	}
	re := regexp.MustCompile(`^local colorscheme = ".*"`)
	res := re.ReplaceAllString(string(csFileContent), "local colorscheme = \""+theme+"\"")

	os.WriteFile(csFileName, []byte(res), 0644)
}

func switchKitty(theme string) {
	kcFileName := viper.GetString("kitty-conf")
	kcFileContent, err := os.ReadFile(kcFileName)
	if err != nil {
		log.Fatal("Could not open kitty-conf")
	}
	re := regexp.MustCompile(`include \./.*\.conf`)
	res := re.ReplaceAllString(string(kcFileContent), "include ./"+theme+".conf")

	os.WriteFile(kcFileName, []byte(res), 0644)
}
