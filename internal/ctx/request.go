package ctx

import "context"

func GetRequestId(ctx context.Context) (string, bool) {
	requestId, ok := ctx.Value(RequestIdKey).(string)
	return requestId, ok
}
