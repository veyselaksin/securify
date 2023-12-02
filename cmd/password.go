/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"crypto/rand"
	"fmt"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

// Flags for the generate command
var (
	digits            *bool
	specialCharacters *bool
	letters           *bool
	capitalLetters    *bool
	passwordLength    *int
)

// global variables for the password command
var (
	charSet string
)

// passwordCmd represents the password command
var passwordCmd = &cobra.Command{
	Use:   "password",
	Short: "Subcommand for generate command",
	Long: `
		password is subcommand for generate command to generate the password for the given arguments
		for example:
		- securify generate password// will generate the password with default values i.e. digits, special characters, letters and capital letters with length 16
		- securify generate password -d -s -l -c // will generate the password with digits, special characters, letters and capital letters with length 16
		- securify generate password -d -s -l -c -n 10 // will generate the password with digits, special characters, letters and capital letters with length 10
		- securify generate password -d -s -n 15 // will generate the password with digits, special characters with length 15
		- securify generate password -d -s -l -c -n 20 // will generate the password with digits, special characters, letters and capital letters with length 20
	`,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: Generate the password
		// TODO: Print the password

		// Check the flags passed in the command line,so if the flag written as --digits or -d then it will be true otherwise false
		setFlagsStatus(cmd)

		charset := getCharSet()
		buffer := make([]byte, *passwordLength)
		_, err := rand.Read(buffer)
		if err != nil {
			color.Red("Error while generating the password")
			return
		}

		charsetLength := len(charset)
		for i := 0; i < *passwordLength; i++ {
			buffer[i] = charset[int(buffer[i])%charsetLength]
		}

		color.Green("Hey buddy, your password is: %s\n", string(buffer))
	},
}

func setFlagsStatus(cmd *cobra.Command) {

	if !isFlagPassed(cmd, "digits") {
		setFlagValue(cmd, "digits", false)
	}
	if !isFlagPassed(cmd, "special-characters") {
		setFlagValue(cmd, "special-characters", false)
	}
	if !isFlagPassed(cmd, "letters") {
		setFlagValue(cmd, "letters", false)
	}
	if !isFlagPassed(cmd, "capital-letters") {
		setFlagValue(cmd, "capital-letters", false)
	}

	// if only length flag is passed then set the default values for the flags
	if isFlagPassed(cmd, "length") && !isFlagPassed(cmd, "digits") && !isFlagPassed(cmd, "special-characters") && !isFlagPassed(cmd, "letters") && !isFlagPassed(cmd, "capital-letters") {
		setFlagValue(cmd, "digits", true)
		setFlagValue(cmd, "special-characters", true)
		setFlagValue(cmd, "letters", true)
		setFlagValue(cmd, "capital-letters", true)
	}
}

func isFlagPassed(cmd *cobra.Command, flagName string) bool {
	return cmd.Flags().Lookup(flagName).Changed
}

func setFlagValue(cmd *cobra.Command, flagName string, value any) {
	flag := cmd.Flags().Lookup(flagName)
	if flag != nil {
		flag.Value.Set(fmt.Sprintf("%v", value))
	}
}

func getCharSet() string {

	if *digits {
		charSet += "0123456789"
	}

	if *specialCharacters {
		charSet += "!@#$%^&*()_+"
	}

	if *letters {
		charSet += "abcdefghijklmnopqrstuvwxyz"
	}

	if *capitalLetters {
		charSet += "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	}

	return charSet
}

func init() {
	generateCmd.AddCommand(passwordCmd)

	// Here you will define your flags and configuration settings.
	digits = passwordCmd.Flags().BoolP("digits", "d", true, "Include digits in the password")
	specialCharacters = passwordCmd.Flags().BoolP("special-characters", "s", true, "Include special characters in the password")
	letters = passwordCmd.Flags().BoolP("letters", "l", true, "Include letters in the password")
	capitalLetters = passwordCmd.Flags().BoolP("capital-letters", "c", true, "Include capital letters in the password")
	passwordLength = passwordCmd.Flags().IntP("length", "n", 16, "Length of the password")

	// if there is no flag passed then set the default values for the flags

}
