/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

// generateCmd represents the generate command
var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "generate command is used to generate the password for the given arguments",
	Long: `
		Generate command is used to generate the password for the given arguments
		For example:
		- securify generate password default // will generate the password with default values i.e. 10 digits, 10 special characters and 10 letters
		- securify generate password -d 10 -s 10 -l 10 // will generate the password with 10 digits, 10 special characters and 10 letters
		- securify generate password -d 10 -s 10 // will generate the password with 10 digits and 10 special characters
		- securify generate password -d 10 // will generate the password with 10 digits
		- securify generate password -s 10 // will generate the password with 10 special characters
		- securify generate password -l 10 // will generate the password with 10 letters
		- securify generate password -d 10 -s 10 -l 10 -c // will generate the password with 10 digits, 10 special characters, 10 letters and 10 capital letters
	`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			if err := cmd.Help(); err != nil {
				color.Red("Error while printing the help")
			}
			return
		}
	},
}

func init() {
	rootCmd.AddCommand(generateCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// generateCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// generateCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
