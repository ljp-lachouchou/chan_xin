package ctxdata

import "context"

func GetUId(ctx context.Context) string {
	if value, ok := ctx.Value(Identify).(string); ok {
		return value
	}
	return ""
}
