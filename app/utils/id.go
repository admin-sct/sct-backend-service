package utils

import (
	"fmt"
	"time"
)

// GenerateID generates a unique ID (simple implementation for demo)
// In production, use UUID or database auto-increment
func GenerateID() string {
	return fmt.Sprintf("%d", time.Now().UnixNano())
}

// IsValidID checks if an ID is valid
func IsValidID(id string) bool {
	return id != ""
}
