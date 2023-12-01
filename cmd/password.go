/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"crypto/rand"
	"fmt"

	"github.com/spf13/cobra"
)

// Flags for the generate command
var (
	digits            *bool
	specialCharacters *bool
	letters           *bool
	capitalLetters    *bool
	passwordLength    *int
	defaultFlag       *bool
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
		- securify generate password// will generate the password with default values i.e. 10 digits, 10 special characters and 10 letters
		- securify generate password -d true -s true -l true // will generate the password with 10 digits, 10 special characters and 10 letters
		- securify generate password -d true -s true // will generate the password with 10 digits and 10 special characters
		- securify generate password -d true // will generate the password with 10 digits
		- securify generate password -s true // will generate the password with 10 special characters
		- securify generate password -l true // will generate the password with 10 letters
		- securify generate password -d true -s true -l true -c true // will generate the password with 10 digits, 10 special characters, 10 letters and 10 capital letters
	`,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: Generate the password
		// TODO: Store the password in the local storage
		// TODO: Print the password

		charset := getCharSet()
		buffer := make([]byte, *passwordLength)
		_, err := rand.Read(buffer)
		if err != nil {
			panic(err)
		}

		charsetLength := len(charset)
		for i := 0; i < *passwordLength; i++ {
			buffer[i] = charset[int(buffer[i])%charsetLength]
		}

		println(string(buffer))
	},
}

// otpChars := "0123456789"
// buffer := make([]byte, length)
// _, err := rand.Read(buffer)
// if err != nil {
// 	return "", err
// }

// otpCharsLength := len(otpChars)
// for i := 0; i < length; i++ {
// 	buffer[i] = otpChars[int(buffer[i])%otpCharsLength]
// }

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

// Initialize the flags for the password command and add it to the generate command
// If the default flag is true then set the default values for the flags else set the values for the flags
// from the command line arguments passed by the user. If the user has not passed any value for the flag then
// the default value of the flag will be used. For example:
// - securify generate password default // will generate the password with default values i.e. 10 digits, 10 special characters and 10 letters
// - securify generate password -d 10 -s 10 -l 10 // will generate the password with 10 digits, 10 special characters and 10 letters
// - securify generate password -d 10 -s 10 // will generate the password with 10 digits and 10 special characters
func init() {
	generateCmd.AddCommand(passwordCmd)

	// Here you will define your flags and configuration settings.
	defaultFlag = passwordCmd.Flags().BoolP("default", "t", false, "Generate the password with default values i.e. 10 digits, 10 special characters and 10 letters")
	digits = passwordCmd.Flags().BoolP("digits", "d", false, "Generate the password with digits")
	specialCharacters = passwordCmd.Flags().BoolP("special-characters", "s", false, "Generate the password with special characters")
	letters = passwordCmd.Flags().BoolP("letters", "l", true, "Generate the password with letters")
	capitalLetters = passwordCmd.Flags().BoolP("capital-letters", "c", false, "Generate the password with capital letters")
	passwordLength = passwordCmd.Flags().IntP("len", "n", 16, "Generate the password with the given length")

	fmt.Println(*defaultFlag)
	if *defaultFlag {
		*digits = true
		*specialCharacters = true
		*letters = true
		*capitalLetters = true
		*passwordLength = 16
	}

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// passwordCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// passwordCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
