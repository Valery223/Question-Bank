package domain

import (
	"context"
)

type ctxKey int

const (
	userIDKey ctxKey = iota
	userRoleKey
)

// СЕТТЕР (исрользует Middleware в слое Delivery)
func NewContextWithUser(ctx context.Context, userID ID, role UserRole) context.Context {
	ctx = context.WithValue(ctx, userIDKey, userID)
	ctx = context.WithValue(ctx, userRoleKey, role)
	return ctx
}

//	ГЕТТЕР (Использует UseCase или Repository)
//
// Достает данные безопасно
func UserFromContext(ctx context.Context) (ID, UserRole, bool) {
	id, okID := ctx.Value(userIDKey).(ID)
	role, okRole := ctx.Value(userRoleKey).(UserRole)

	if !okID || !okRole {
		return ID(""), RoleNone, false
	}
	return id, role, true
}
