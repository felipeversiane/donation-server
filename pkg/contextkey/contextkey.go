package contextkey

// Key defines a context key type for request-scoped values.
type Key string

const (
	RequestID Key = "request_id"
	UserID    Key = "user_id"
)
