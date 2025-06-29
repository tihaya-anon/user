package context_key

import "context"

func withKey(ctx context.Context, key any, value any) context.Context {
	return context.WithValue(ctx, key, value)
}

func getKey[T any](ctx context.Context, key any) T {
	if v, ok := ctx.Value(key).(T); ok {
		return v
	}
	return *new(T)
}
