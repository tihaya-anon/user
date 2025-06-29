package context_key

import "context"

type idempotencyKeyKey struct{}

func WithIdempotencyKey(ctx context.Context, idempotencyKey string) context.Context {
	return withKey(ctx, idempotencyKeyKey{}, idempotencyKey)
}

func GetIdempotencyKey(ctx context.Context) string {
	return getKey[string](ctx, idempotencyKeyKey{})
}
