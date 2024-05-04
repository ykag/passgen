package cmd

import (
	"strings"
	"testing"
)

func TestPasswordLength(t *testing.T) {
	testCases := []struct {
		complexity     string
		expectedLength int
	}{
		{"low", 12},
		{"medium", 16},
		{"high", 24},
	}

	for _, tc := range testCases {
		t.Run(tc.complexity, func(t *testing.T) {
			password := generatePassword(tc.complexity, 16)

			if len(password) != tc.expectedLength {
				t.Errorf("Expected password length %d for complexity %s, but got length %d", tc.expectedLength, tc.complexity, len(password))
			}

			if tc.complexity == "low" && (!containsAny(password, lowerChars) || !containsAny(password, upperChars) || !containsAny(password, digitChars)) {
				t.Errorf("Low complexity password does not contain required character types")
			}

			if tc.complexity == "medium" && (!containsAny(password, lowerChars) || !containsAny(password, upperChars) || !containsAny(password, digitChars) || !containsAny(password, symbolChars)) {
				t.Errorf("Medium complexity password does not contain required character types")
			}

			if tc.complexity == "high" && (!containsAny(password, lowerChars) || !containsAny(password, upperChars) || !containsAny(password, digitChars) || !containsAny(password, symbolChars)) {
				t.Errorf("High complexity password does not contain required character types")
			}
		})
	}
}

func containsAny(password, charset string) bool {
	for _, char := range charset {
		if strings.ContainsRune(password, char) {
			return true
		}
	}

	return false
}
