package logger

import (
	"context"
	"testing"
)

func TestContextWithRequestIDAndUserID(t *testing.T) {
	ctx := context.Background()
	ctx = ContextWithRequestID(ctx, "req123")
	ctx = ContextWithUserID(ctx, "user456")

	rid, ok := RequestIDFromContext(ctx)
	if !ok || rid != "req123" {
		t.Fatalf("RequestIDFromContext=%q, ok=%v", rid, ok)
	}
	uid, ok := UserIDFromContext(ctx)
	if !ok || uid != "user456" {
		t.Fatalf("UserIDFromContext=%q, ok=%v", uid, ok)
	}
}
