/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "securify",
	Short: "Generate a new password with a simple command line interface",
	Long: `
		┌──────────────────────────────────────────────────────────┐
		│                                                          │
		│       ____                            _   __             │
		│      / ___|   ___   ___  _   _  _ __ (_) / _| _   _      │
		│      \___ \  / _ \ / __|| | | || '__|| || |_ | | | |     │
		│       ___) ||  __/| (__ | |_| || |   | ||  _|| |_| |     │
		│      |____/  \___| \___| \__,_||_|   |_||_|   \__, |     │
		│                                               |___/      │
		│                                                          │
		└──────────────────────────────────────────────────────────┘
		Version: 0.0.1
		Author:  Veysel Aksin

		Create a new password with a simple command line interface

		Usage:
		securify [command] [flags] [arguments]

		Available Commands:
		generate    Generate a new password
		help        Help about any command
	`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.securify.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
