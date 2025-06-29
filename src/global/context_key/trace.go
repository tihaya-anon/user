package context_key

import "context"

type traceIdKey struct{}

func WithTraceId(ctx context.Context, traceId string) context.Context {
	return withKey(ctx, traceIdKey{}, traceId)
}
func GetTraceId(ctx context.Context) string {
	return getKey[string](ctx, traceIdKey{})
}
