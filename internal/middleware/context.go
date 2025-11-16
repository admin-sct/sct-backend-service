package middleware

import (
	"context"
)

// UpdateContext updates the context with additional values
// This can be used to add tracing, user info, etc.
func UpdateContext(ctx context.Context) context.Context {
	// TODO: Add tracing context, user info, etc.
	return ctx
}

// GetUserID retrieves user ID from context
func GetUserID(ctx context.Context) (string, bool) {
	// TODO: Extract user ID from context
	userID, ok := ctx.Value("user_id").(string)
	return userID, ok
}

// WithUserID adds user ID to context
func WithUserID(ctx context.Context, userID string) context.Context {
	return context.WithValue(ctx, "user_id", userID)
}

