package context_key

import "context"

type requestIdKey struct{}

func WithRequestId(ctx context.Context, requestId string) context.Context {
	return withKey(ctx, requestIdKey{}, requestId)
}

func GetRequestId(ctx context.Context) string {
	return getKey[string](ctx, requestIdKey{})
}
