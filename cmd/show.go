/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

// showCmd represents the show command
var showCmd = &cobra.Command{
	Use:   "show",
	Short: "Show the all the passwords",
	Long: `
		Show the all the passwords in securify/passwords.spf file which is located in the home directory
		For example:
		- securify show
`,
	Run: func(cmd *cobra.Command, args []string) {
		// Read the password file
		readPasswordFile()
	},
}

func init() {
	rootCmd.AddCommand(showCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// showCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// showCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// Read the password file
func readPasswordFile() {
	data, err := os.ReadFile(passwordFilePath)
	if err != nil {
		return
	}

	// Print the data
	fmt.Println(string(data))
}
