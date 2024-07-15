package cmd

import (
	"testing"
)

func TestGeneratePassword(t *testing.T) {
	t.Run("should generate the password with default values", func(t *testing.T) {

		// Act
		password, err := generatePassword()
		if err != nil {
			t.Errorf("error should be nil")
		}

		if len(password) != 16 {
			t.Errorf("password length should be 16")
		}
	})

	t.Run("should generate the password only with digits", func(t *testing.T) {

		// Arrange
		*passwordLength = 10
		*digits = true
		*specialCharacters = false
		*letters = false
		*capitalLetters = false

		// Act
		password, err := generatePassword()
		if err != nil {
			t.Errorf("error should be nil")
		}

		if len(password) != 10 {
			t.Errorf("password length should be 10")
		}
	})

	t.Run("should generate the password only with special characters", func(t *testing.T) {

		// Arrange
		*passwordLength = 10
		*digits = false
		*specialCharacters = true
		*letters = false
		*capitalLetters = false

		// Act
		password, err := generatePassword()
		if err != nil {
			t.Errorf("error should be nil")
		}

		if len(password) != 10 {
			t.Errorf("password length should be 10")
		}
	})

	t.Run("should generate the password only with letters", func(t *testing.T) {

		// Arrange
		*passwordLength = 10
		*digits = false
		*specialCharacters = false
		*letters = true
		*capitalLetters = false

		// Act
		password, err := generatePassword()
		if err != nil {
			t.Errorf("error should be nil")
		}

		if len(password) != 10 {
			t.Errorf("password length should be 10")
		}
	})

	t.Run("should generate the password only with capital letters", func(t *testing.T) {

		// Arrange
		*passwordLength = 10
		*digits = false
		*specialCharacters = false
		*letters = false
		*capitalLetters = true

		// Act
		password, err := generatePassword()
		if err != nil {
			t.Errorf("error should be nil")
		}

		if len(password) != 10 {
			t.Errorf("password length should be 10")
		}
	})

	t.Run("should generate the password with digits and special characters", func(t *testing.T) {

		// Arrange
		*passwordLength = 10
		*digits = true
		*specialCharacters = true
		*letters = false
		*capitalLetters = false

		// Act
		password, err := generatePassword()
		if err != nil {
			t.Errorf("error should be nil")
		}

		if len(password) != 10 {
			t.Errorf("password length should be 10")
		}
	})

	t.Run("should generate the password with digits and letters", func(t *testing.T) {

		// Arrange
		*passwordLength = 10
		*digits = true
		*specialCharacters = false
		*letters = true
		*capitalLetters = false

		// Act
		password, err := generatePassword()
		if err != nil {
			t.Errorf("error should be nil")
		}

		if len(password) != 10 {
			t.Errorf("password length should be 10")
		}
	})

	t.Run("should generate the password with digits and capital letters", func(t *testing.T) {

		// Arrange
		*passwordLength = 10
		*digits = true
		*specialCharacters = false
		*letters = false
		*capitalLetters = true

		// Act
		password, err := generatePassword()
		if err != nil {
			t.Errorf("error should be nil")
		}

		if len(password) != 10 {
			t.Errorf("password length should be 10")
		}
	})

	t.Run("should generate the password with special characters and letters", func(t *testing.T) {

		// Arrange
		*passwordLength = 10
		*digits = false
		*specialCharacters = true
		*letters = true
		*capitalLetters = false

		// Act
		password, err := generatePassword()
		if err != nil {
			t.Errorf("error should be nil")
		}

		if len(password) != 10 {
			t.Errorf("password length should be 10")
		}
	})

	t.Run("should generate the password with special characters and capital letters", func(t *testing.T) {

		// Arrange
		*passwordLength = 10
		*digits = false
		*specialCharacters = true
		*letters = false
		*capitalLetters = true

		// Act
		password, err := generatePassword()
		if err != nil {
			t.Errorf("error should be nil")
		}

		if len(password) != 10 {
			t.Errorf("password length should be 10")
		}
	})

	t.Run("should generate the password with letters and capital letters", func(t *testing.T) {

		// Arrange
		*passwordLength = 10
		*digits = false
		*specialCharacters = false
		*letters = true
		*capitalLetters = true

		// Act
		password, err := generatePassword()
		if err != nil {
			t.Errorf("error should be nil")
		}

		if len(password) != 10 {
			t.Errorf("password length should be 10")
		}

	})

	t.Run("should generate the password with digits, special characters and letters", func(t *testing.T) {

		// Arrange
		*passwordLength = 10
		*digits = true
		*specialCharacters = true
		*letters = true
		*capitalLetters = false

		// Act
		password, err := generatePassword()
		if err != nil {
			t.Errorf("error should be nil")
		}

		if len(password) != 10 {
			t.Errorf("password length should be 10")
		}

	})

	t.Run("should generate the password with digits, special characters and capital letters", func(t *testing.T) {

		// Arrange
		*passwordLength = 10
		*digits = true
		*specialCharacters = true
		*letters = false
		*capitalLetters = true

		// Act
		password, err := generatePassword()
		if err != nil {
			t.Errorf("error should be nil")
		}

		if len(password) != 10 {
			t.Errorf("password length should be 10")
		}

	})

	t.Run("should generate the password with digits, letters and capital letters", func(t *testing.T) {

		// Arrange
		*passwordLength = 10
		*digits = true
		*specialCharacters = false
		*letters = true
		*capitalLetters = true

		// Act
		password, err := generatePassword()
		if err != nil {
			t.Errorf("error should be nil")
		}

		if len(password) != 10 {
			t.Errorf("password length should be 10")
		}

	})

	t.Run("should generate the password with special characters, letters and capital letters", func(t *testing.T) {

		// Arrange
		*passwordLength = 10
		*digits = false
		*specialCharacters = true
		*letters = true
		*capitalLetters = true

		// Act
		password, err := generatePassword()
		if err != nil {
			t.Errorf("error should be nil")
		}

		if len(password) != 10 {
			t.Errorf("password length should be 10")
		}

	})

	t.Run("should generate the password with digits, special characters, letters and capital letters", func(t *testing.T) {

		// Arrange
		*passwordLength = 10
		*digits = true
		*specialCharacters = true
		*letters = true
		*capitalLetters = true

		// Act
		password, err := generatePassword()
		if err != nil {
			t.Errorf("error should be nil")
		}

		if len(password) != 10 {
			t.Errorf("password length should be 10")
		}

	})

	t.Run("should generate the password with digits, special characters, letters and capital letters with length 20", func(t *testing.T) {

		// Arrange
		*passwordLength = 20
		*digits = true
		*specialCharacters = true
		*letters = true
		*capitalLetters = true

		// Act
		password, err := generatePassword()
		if err != nil {
			t.Errorf("error should be nil")
		}

		if len(password) != 20 {
			t.Errorf("password length should be 20")
		}

	})

	t.Run("should generate the password with only length flag", func(t *testing.T) {

		// Arrange
		*passwordLength = 20
		*digits = false
		*specialCharacters = false
		*letters = false
		*capitalLetters = false

		// Act
		password, err := generatePassword()
		if err != nil {
			t.Errorf("error should be nil")
		}

		if len(password) != 20 {
			t.Errorf("password length should be 20")
		}

	})

	t.Run("should generate the password with only length flag and digits", func(t *testing.T) {

		// Arrange
		*passwordLength = 20
		*digits = true
		*specialCharacters = false
		*letters = false
		*capitalLetters = false

		// Act
		password, err := generatePassword()
		if err != nil {
			t.Errorf("error should be nil")
		}

		if len(password) != 20 {
			t.Errorf("password length should be 20")
		}

	})

	t.Run("should generate the password with only length flag and special characters", func(t *testing.T) {

		// Arrange
		*passwordLength = 20
		*digits = false
		*specialCharacters = true
		*letters = false
		*capitalLetters = false

		// Act
		password, err := generatePassword()
		if err != nil {
			t.Errorf("error should be nil")
		}

		if len(password) != 20 {
			t.Errorf("password length should be 20")
		}

	})

	t.Run("should generate the password with only length flag and letters", func(t *testing.T) {

		// Arrange
		*passwordLength = 20
		*digits = false
		*specialCharacters = false
		*letters = true
		*capitalLetters = false

		// Act
		password, err := generatePassword()
		if err != nil {
			t.Errorf("error should be nil")
		}

		if len(password) != 20 {
			t.Errorf("password length should be 20")
		}

	})

	t.Run("should generate the password with only length flag and capital letters", func(t *testing.T) {

		// Arrange
		*passwordLength = 20
		*digits = false
		*specialCharacters = false
		*letters = false
		*capitalLetters = true

		// Act
		password, err := generatePassword()
		if err != nil {
			t.Errorf("error should be nil")
		}

		if len(password) != 20 {
			t.Errorf("password length should be 20")
		}

	})

	t.Run("should generate the password with only length flag and digits, special characters", func(t *testing.T) {

		// Arrange
		*passwordLength = 20
		*digits = true
		*specialCharacters = true
		*letters = false
		*capitalLetters = false

		// Act
		password, err := generatePassword()
		if err != nil {
			t.Errorf("error should be nil")
		}

		if len(password) != 20 {
			t.Errorf("password length should be 20")
		}

	})

	t.Run("should error if the length is less than or equal 0", func(t *testing.T) {

		// Arrange
		*passwordLength = 0
		*digits = false
		*specialCharacters = false
		*letters = false
		*capitalLetters = false

		// Act
		password, err := generatePassword()
		if err == nil {
			t.Errorf("error should not be nil")
		}

		if password != "" {
			t.Errorf("password should be empty")
		}
	})

	t.Run("should error if the length is greater than or equal 100", func(t *testing.T) {

		// Arrange
		*passwordLength = 101
		*digits = false
		*specialCharacters = false
		*letters = false
		*capitalLetters = false

		// Act
		password, err := generatePassword()
		if err == nil {
			t.Errorf("error should not be nil")
		}

		if password != "" {
			t.Errorf("password should be empty")
		}
	})

}
