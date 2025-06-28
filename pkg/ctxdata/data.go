package ctxdata

import "context"

func GetUId(ctx context.Context) string {
	if value, ok := ctx.Value(Identify).(string); ok {
		return value
	}
	return ""
}
func GetParams(ctx context.Context, identify string) any {
	if value := ctx.Value(identify); value != nil {
		return value
	}
	return nil
}
