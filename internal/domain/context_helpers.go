package domain

import (
	"context"
)

// Ключи для значений в контексте
type ctxKey int

const (
	userIDKey ctxKey = iota
	userRoleKey
)

// NEWTER (Используется в Middleware)
//
// Создает новый контекст с информацией о пользователе
func NewContextWithUser(ctx context.Context, userID ID, role UserRole) context.Context {
	ctx = context.WithValue(ctx, userIDKey, userID)
	ctx = context.WithValue(ctx, userRoleKey, role)
	return ctx
}

// UserFromContext извлекает информацию о пользователе из контекста
func UserFromContext(ctx context.Context) (ID, UserRole, bool) {
	id, okID := ctx.Value(userIDKey).(ID)
	role, okRole := ctx.Value(userRoleKey).(UserRole)

	if !okID || !okRole {
		return ID(""), RoleNone, false
	}
	return id, role, true
}
