package utils

import (
	"regexp"
)

var (
	emailRegex = regexp.MustCompile(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`)
)

// IsValidEmail validates an email address
func IsValidEmail(email string) bool {
	if email == "" {
		return false
	}
	return emailRegex.MatchString(email)
}

// IsValidString checks if a string is not empty
func IsValidString(s string) bool {
	return len(s) > 0
}

// TruncateString truncates a string to a maximum length
func TruncateString(s string, maxLen int) string {
	if len(s) <= maxLen {
		return s
	}
	return s[:maxLen]
}
