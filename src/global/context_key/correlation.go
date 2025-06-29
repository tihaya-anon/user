package context_key

import "context"

type correlationIdKey struct{}

func WithCorrelationId(ctx context.Context, correlationId string) context.Context {
	return withKey(ctx, correlationIdKey{}, correlationId)
}

func GetCorrelationId(ctx context.Context) string {
	return getKey[string](ctx, correlationIdKey{})
}
