package cmd

import (
	"os"
	"path/filepath"
)

var homeDir string
var securifyDir string
var passwordFilePath string

func init() {
	var err error
	homeDir, err = os.UserHomeDir()
	if err != nil {
		panic(err)
	}

	// Securify is a folder in the home directory
	securifyDir = filepath.Join(homeDir, "securify")
	_, err = os.Stat(securifyDir)
	if err != nil {
		// Create the securify directory
		err := os.Mkdir(securifyDir, 0755)
		if err != nil {
			panic(err)
		}
	}

	// Create the password file
	_, err = os.Stat(filepath.Join(securifyDir, "password.spf"))
	if err != nil {
		// Create the password file
		_, err := os.Create(filepath.Join(securifyDir, "password.spf"))
		if err != nil {
			panic(err)
		}
	}

	passwordFilePath = filepath.Join(securifyDir, "password.spf")
}
