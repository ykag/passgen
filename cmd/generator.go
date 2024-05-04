package cmd

import (
	"fmt"
	"math/rand"
	"os"

	"github.com/spf13/cobra"
)

const (
	lowerChars   = "abcdefghijklmnopqrstuvwxyz"
	upperChars   = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	digitChars   = "0123456789"
	symbolChars  = "!@#$%^&*()-_=+,.?"
	defaultChars = lowerChars + upperChars + digitChars + symbolChars
)

func init() {
	generateCmd.Flags().String("complexity", "medium", "Complexity level (low, medium, high)")
	generateCmd.Flags().Int("length", 16, "Password length")

	rootCmd.AddCommand(generateCmd)
}

var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generate a new random password",
	Long:  `Generate a new random password`,
	Run: func(cmd *cobra.Command, args []string) {
		complexity, _ := cmd.Flags().GetString("complexity")
		length, _ := cmd.Flags().GetInt("length")

		if length <= 0 {
			fmt.Println("Password length must be greater than 0")
			os.Exit(1)
		}

		password := generatePassword(complexity, length)
		fmt.Println(password)
	},
}

func generatePassword(complexity string, length int) string {
	var chars string
	switch complexity {
	case "low":
		length = 12
		chars = lowerChars + upperChars + digitChars
	case "medium":
		chars = lowerChars + upperChars + digitChars + symbolChars
	case "high":
		length = 24
		chars = defaultChars
	default:
		fmt.Println("Invalid complexity level. Please choose low, medium or high.")
		os.Exit(1)
	}

	return generateRandomPassword(chars, length)
}

func generateRandomPassword(chars string, length int) string {
	password := make([]byte, length)
	for i := range password {
		password[i] = chars[rand.Intn(len(chars))]
	}

	return string(password)
}
