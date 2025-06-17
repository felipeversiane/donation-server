package logger

import "context"

// contextKey is an unexported type for keys defined in this package.
type contextKey string

const (
	// RequestIDKey stores a unique request identifier in the context.
	RequestIDKey contextKey = "request_id"
	// UserIDKey stores an authenticated user identifier in the context.
	UserIDKey contextKey = "user_id"
)

// ContextWithRequestID returns a new context with the request id set.
func ContextWithRequestID(ctx context.Context, id string) context.Context {
	return context.WithValue(ctx, RequestIDKey, id)
}

// ContextWithUserID returns a new context with the user id set.
func ContextWithUserID(ctx context.Context, userID string) context.Context {
	return context.WithValue(ctx, UserIDKey, userID)
}

// RequestIDFromContext retrieves the request id from the context.
func RequestIDFromContext(ctx context.Context) (string, bool) {
	rid, ok := ctx.Value(RequestIDKey).(string)
	return rid, ok
}

// UserIDFromContext retrieves the user id from the context.
func UserIDFromContext(ctx context.Context) (string, bool) {
	uid, ok := ctx.Value(UserIDKey).(string)
	return uid, ok
}
