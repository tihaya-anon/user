package context_key

import "context"

type traceIdKey struct{}

func WithTraceId(ctx context.Context, traceId string) context.Context {
	return context.WithValue(ctx, traceIdKey{}, traceId)
}
func GetTraceId(ctx context.Context) string {
	if v, ok := ctx.Value(traceIdKey{}).(string); ok {
		return v
	}
	return ""
}
