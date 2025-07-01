package context_key

import "context"

type jwtKey struct{}

func WithJwt(ctx context.Context, jwt string) context.Context {
	return withKey(ctx, jwtKey{}, jwt)
}

func GetJwt(ctx context.Context) string {
	return getKey[string](ctx, jwtKey{})
}
