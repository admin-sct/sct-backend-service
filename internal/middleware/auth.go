package middleware

import (
	"context"
	"net/http"
)

// AuthMiddleware handles authentication
func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// TODO: Implement authentication logic
		// Extract token, validate, add user info to context
		ctx := r.Context()

		// For now, just pass through
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// RequireAuth ensures the request is authenticated
func RequireAuth(ctx context.Context) error {
	// TODO: Check if user is authenticated
	_, ok := GetUserID(ctx)
	if !ok {
		return ErrUnauthenticated
	}
	return nil
}

var (
	ErrUnauthenticated = &AuthError{Message: "unauthorized"}
)

// AuthError represents an authentication error
type AuthError struct {
	Message string
}

func (e *AuthError) Error() string {
	return e.Message
}

