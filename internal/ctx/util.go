package ctx

import "context"

type ctxKey int

const (
	_ ctxKey = iota
	ctxKeyRequestID
	ctxKeyDB
)

func WithRequestID(ctx context.Context, requestID string) context.Context {
	return context.WithValue(ctx, ctxKeyRequestID, requestID)
}

func RequestID(ctx context.Context) string {
	requestID := ctx.Value(ctxKeyRequestID)
	if requestID != nil {
		return requestID.(string)
	}

	return ""
}

func WithDB(ctx context.Context, db string) context.Context {
	return context.WithValue(ctx, ctxKeyDB, db)
}

func DB(ctx context.Context) string {
	db := ctx.Value(ctxKeyDB)
	if db != nil {
		return db.(string)
	}

	return ""
}
