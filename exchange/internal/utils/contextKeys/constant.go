package contextKeys

import (
	"context"

	"exchange/internal/enum/permission"
	"pkg/errors"
)

type ContextKey int

const (
	permissionsKey ContextKey = iota + 1
	isRootKey
	uuidKey
	XRequestIDKey
)

func SetIsRoot(ctx context.Context, isRoot bool) context.Context {
	return context.WithValue(ctx, isRootKey, isRoot)
}

func SetPermissions(ctx context.Context, permissions []permission.Permission) context.Context {
	return context.WithValue(ctx, permissionsKey, permissions)
}

func SetUUID(ctx context.Context, uuid string) context.Context {
	return context.WithValue(ctx, uuidKey, uuid)
}

func GetUUID(ctx context.Context) (string, error) {
	uuid, ok := ctx.Value(uuidKey).(string)

	if !ok {
		return "", errors.BadRequest.New("Uuid was not found")
	}

	return uuid, nil
}

func GetXRequestID(ctx context.Context) (string, error) {
	xRequestID, ok := ctx.Value(XRequestIDKey).(string)

	if !ok {
		return "", errors.BadRequest.New("XRequestID was not found")
	}

	return xRequestID, nil
}
