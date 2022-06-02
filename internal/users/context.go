package users

import (
	"context"
)

func GetCtxStringVal(ctx context.Context, key ContextKey) string {
	ctxValue := ctx.Value(key)

	if ctxValue != nil {
		val, ok := ctxValue.(string)
		if ok {
			return val
		}
	}

	return ""
}
