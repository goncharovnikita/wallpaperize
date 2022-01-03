package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// App is application interface which will be used
// in all commands
type App interface {
	Daily() error
	Info(format string) error
	Random(loadOnly bool) error
	Restore() error
	Set(path string) error
	GetSelected() error
	Place() error
}

// Execute starts programm
func Execute(app App) {
	rootCmd := createRootCommand()

	rootCmd.AddCommand(createSetCommand(app))
	rootCmd.AddCommand(createDailyCommand(app))
	rootCmd.AddCommand(createGetCommand(app))
	rootCmd.AddCommand(createGetSelectedCommand(app))
	rootCmd.AddCommand(createInfoCommand(app))
	rootCmd.AddCommand(createPlaceCommand(app))
	rootCmd.AddCommand(createRandomCommand(app))
	rootCmd.AddCommand(createRestoreCommand(app))

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

// rootCmd is app entry point
func createRootCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "wallpaperize",
		Short: "Wallpaperize is tool for setting up wallpaper quickly",
		Run: func(cmd *cobra.Command, args []string) {
			if err := cmd.Usage(); err != nil {
				fmt.Errorf(err.Error())
			}
		},
	}
}

// setCmd represents the set command
func createSetCommand(app App) *cobra.Command {
	return &cobra.Command{
		Use:   "set",
		Short: "Set image from given path as wallpaper",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) < 1 {
				fmt.Errorf("path to file must be specified")
			} else {
				if err := app.Set(args[0]); err != nil {
					fmt.Errorf(err.Error())
				}
			}
		},
	}
}

// dailyCmd represents the daily command
func createDailyCommand(app App) *cobra.Command {
	return &cobra.Command{
		Use:   "daily",
		Short: "Sets daily image as wallpaper",
		Run: func(cmd *cobra.Command, args []string) {
			if err := app.Daily(); err != nil {
				fmt.Errorf(err.Error())
			}
		},
	}
}

func createGetCommand(app App) *cobra.Command {
	return &cobra.Command{
		Use:   "get",
		Short: "Get something that wallpaperize can give",
		Run: func(cmd *cobra.Command, args []string) {
			if err := cmd.Usage(); err != nil {
				fmt.Errorf(err.Error())
			}
		},
	}
}

func createGetSelectedCommand(app App) *cobra.Command {
	return &cobra.Command{
		Use:   "selected",
		Short: "Get selected wallpaper path",
		Run: func(cmd *cobra.Command, args []string) {
			if err := app.GetSelected(); err != nil {
				fmt.Errorf(err.Error())
			}
		},
	}
}

var infoOutputType string

func createInfoCommand(app App) *cobra.Command {
	return &cobra.Command{
		Use:   "info",
		Short: "Get info about wallpaperize disk usage",
		Run: func(cmd *cobra.Command, args []string) {
			if err := app.Info(infoOutputType); err != nil {
				fmt.Errorf(err.Error())
			}
		},
	}
}

func createPlaceCommand(app App) *cobra.Command {
	return &cobra.Command{
		Use:   "place",
		Short: "Place print wallpaperize bin placement",
		Run: func(cmd *cobra.Command, args []string) {
			if err := app.Place(); err != nil {
				fmt.Errorf(err.Error())
			}
		},
	}
}

var randomLoadOnly bool

func createRandomCommand(app App) *cobra.Command {
	return &cobra.Command{
		Use:   "random",
		Short: "Set random image from internet as wallpaper",
		Run: func(cmd *cobra.Command, args []string) {
			if err := app.Random(randomLoadOnly); err != nil {
				fmt.Errorf(err.Error())
			}
		},
	}
}

func createRestoreCommand(app App) *cobra.Command {
	return &cobra.Command{
		Use:   "restore",
		Short: "Set initial desktop wallpaper",
		Run: func(cmd *cobra.Command, args []string) {
			if err := app.Restore(); err != nil {
				fmt.Errorf(err.Error())
			}
		},
	}
}
