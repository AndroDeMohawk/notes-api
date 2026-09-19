package Ctx

import "context"

type ctxKey string

const Key = "ctx"

func FromContext(ctx context.Context) string {
	if id, ok := ctx.Value(Key).(string); ok {
		return id
	}
	return ""

}
