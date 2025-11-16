package middleware

import (
	"encoding/json"
	"log"
	"net/http"
)

// ErrorHandler handles errors in HTTP requests
func ErrorHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				log.Printf("Panic recovered: %v", err)
				WriteErrorResponse(w, http.StatusInternalServerError, "Internal server error")
			}
		}()

		next.ServeHTTP(w, r)
	})
}

// WriteErrorResponse writes an error response in JSON format
func WriteErrorResponse(w http.ResponseWriter, statusCode int, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	response := map[string]interface{}{
		"error": map[string]interface{}{
			"message": message,
			"code":    statusCode,
		},
	}

	if err := json.NewEncoder(w).Encode(response); err != nil {
		log.Printf("Failed to encode error response: %v", err)
	}
}

// WriteGraphQLError writes a GraphQL error response
func WriteGraphQLError(w http.ResponseWriter, errors []error) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	response := map[string]interface{}{
		"errors": errors,
	}

	if err := json.NewEncoder(w).Encode(response); err != nil {
		log.Printf("Failed to encode GraphQL error response: %v", err)
	}
}

