package cmd

import (
	"crypto/rand"
	"fmt"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"os"
	"path/filepath"
	"time"
)

// Flags for the generate command
var (
	digits            *bool
	specialCharacters *bool
	lower             *bool
	upper             *bool
	passwordLength    *int
	name              *string
	description       *string
)

// passwordCmd represents the password command
var passwordCmd = &cobra.Command{
	Use:   "pwd",
	Short: "Generate the password for the given arguments",
	Long: `
		Generate the password for the given arguments,
		for example:
		- securify pwd // generates a password with default values (digits, special characters, lowercase and uppercase letters) with length 16
		- securify pwd -d -s -l -u // generates a password with digits, special characters, lowercase and uppercase letters with length 16
		- securify pwd -d -s -l -u --len 10 // generates a password with digits, special characters, lowercase and uppercase letters with length 10
		- securify pwd -d -s --len 15 // generates a password with digits, special characters with length 15
		- securify pwd -d -s -l -u --len 20 // generates a password with digits, special characters, lowercase and uppercase letters with length 20
`,
	Run: func(cmd *cobra.Command, args []string) {
		if isFlagPassed(cmd, "len") && *passwordLength <= 8 {
			color.Red("Password length should be greater than 8")
			return
		}

		// Check which flags are passed except the password length, name and description
		if isFlagPassed(cmd, "len") || isFlagPassed(cmd, "name") || isFlagPassed(cmd, "desc") {
			if !isFlagPassed(cmd, "digits") && !isFlagPassed(cmd, "special") && !isFlagPassed(cmd, "lower") && !isFlagPassed(cmd, "upper") {
				setFlagValue(cmd, "digits", true)
				setFlagValue(cmd, "special", true)
				setFlagValue(cmd, "lower", true)
				setFlagValue(cmd, "upper", true)
			}
		}

		if isFlagPassed(cmd, "digits") {
			setFlagValue(cmd, "digits", true)
		}

		if isFlagPassed(cmd, "special") {
			setFlagValue(cmd, "special", true)
		}

		if isFlagPassed(cmd, "lower") {
			setFlagValue(cmd, "lower", true)
		}

		if isFlagPassed(cmd, "upper") {
			setFlagValue(cmd, "upper", true)
		}

		password, err := generatePassword()
		if err != nil {
			color.Red("Error while generating the password")
			return
		}

		if password == "" {
			return
		}
		color.Green("Your password is: %s\n", password)

		// Save the password to the file if the name is provided
		if *name != "" {
			savePasswordToFile(password, *name, *description)
			color.Green("Password saved successfully")
		}
	},
}

func generatePassword() (string, error) {
	if *passwordLength <= 0 {
		color.Red("Password length should be greater than 0")
		return "", nil
	}

	if *passwordLength > 100 {
		color.Red("Password length should be less than 100")
		return "", nil
	}

	charset := getCharSet()
	buffer := make([]byte, *passwordLength)
	_, err := rand.Read(buffer)
	if err != nil {
		color.Red("Error while generating the password")
		return "", err
	}

	charsetLength := len(charset)
	for i := 0; i < *passwordLength; i++ {
		buffer[i] = charset[int(buffer[i])%charsetLength]
	}

	return string(buffer), nil
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
	var charSet string

	if *digits {
		charSet += "0123456789"
	}

	if *specialCharacters {
		charSet += "!@#$%^&*()_+"
	}

	if *lower {
		charSet += "abcdefghijklmnopqrstuvwxyz"
	}

	if *upper {
		charSet += "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	}

	return charSet
}

func init() {
	rootCmd.AddCommand(passwordCmd)

	// Define flags and configuration settings
	digits = passwordCmd.Flags().BoolP("digits", "d", false, "Include digits in the password")
	specialCharacters = passwordCmd.Flags().BoolP("special", "s", false, "Include special characters in the password")
	lower = passwordCmd.Flags().BoolP("lower", "l", false, "Include lowercase letters in the password")
	upper = passwordCmd.Flags().BoolP("upper", "u", false, "Include uppercase letters in the password")
	passwordLength = passwordCmd.Flags().IntP("len", "", 16, "Length of the password")

	name = passwordCmd.Flags().StringP("save", "", "", "Name of the password")
	description = passwordCmd.Flags().StringP("desc", "C", "", "Description of the password")

}

// Path: homeDir/securify/password.spf
// SPF: Secure Password File
// File Format:
// [password name] [password] [date-time] [description]
// Example:
// [myPassword] [myPassword123] [2023-07-01 12:00:00] [This is my password]

func savePasswordToFile(password string, passName string, description string) {
	securifyDir := filepath.Join(os.Getenv("HOME"), "securify")
	if _, err := os.Stat(securifyDir); os.IsNotExist(err) {
		os.MkdirAll(securifyDir, 0755)
	}

	file, err := os.OpenFile(passwordFilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		color.Red("Error while opening the password file")
		return
	}
	defer file.Close()

	_, err = file.WriteString(fmt.Sprintf("[%s] [%s] [%s] [%s]\n", passName, password, time.Now().Format("2006-01-02 15:04:05"), description))
	if err != nil {
		color.Red("Error while writing the password to file")
	}
}
