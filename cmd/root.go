package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var app App

// rootCmd is app entry point
var rootCmd = &cobra.Command{
	Use:   "wallpaperize",
	Short: "Wallpaperize is tool for setting up wallpaper quickly",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Usage()
	},
}

// Execute starts programm
func Execute(a App) {
	app = a
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(dailyCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// dailyCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// dailyCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
