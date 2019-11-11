package dcontext

import (
	// https://qiita.com/marnie_ms4/items/985d67c4c1b29e11fffc
	"context"
)

type key string

const (
	userIDKey key = "userID"
)

func SetUserID(ctx context.Context, userID string) context.Context {
	return context.WithValue(ctx, userIDKey, userID)
}

func GetUserIDFromContext(ctx context.Context) string {
	var userID string
	if ctx.Value(userIDKey) != nil {
		userID = ctx.Value(userIDKey).(string)
	}
	return userID
}
