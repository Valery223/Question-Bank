package domain

import "time"

type TestSession struct {
	ID         string
	TemolateID string
	UserId     string
	StartedAt  time.Time
	ExpiredAt  time.Time

	// Список вопросов именно для этой попытки
	// (если захотим перемешать вопросы)
	// Почему не id?
	// Так как вопросы могут меняться в шаблоне,
	// а сессия должна хранить именно те вопросы,
	// которые были на момент старта сессии
	Questions []Question
}

// Разные проверки
func (s *TestSession) IsExpired() bool {
	// Логика: если  ExpiredAt в прошлом, то сессия истекла
	return s.ExpiredAt.Before(time.Now())
}
