package logger

import "context"

type contextKey string

const (
	RequestIDKey contextKey = "request_id"
	UserIDKey contextKey = "user_id"
)

func ContextWithRequestID(ctx context.Context, id string) context.Context {
	return context.WithValue(ctx, RequestIDKey, id)
}

func ContextWithUserID(ctx context.Context, userID string) context.Context {
	return context.WithValue(ctx, UserIDKey, userID)
}

func RequestIDFromContext(ctx context.Context) (string, bool) {
	rid, ok := ctx.Value(RequestIDKey).(string)
	return rid, ok
}

func UserIDFromContext(ctx context.Context) (string, bool) {
	uid, ok := ctx.Value(UserIDKey).(string)
	return uid, ok
}
