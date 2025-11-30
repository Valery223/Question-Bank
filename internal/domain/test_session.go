package domain

import "time"

// TestSession представляет сессию тестирования пользователя
type TestSession struct {
	ID         ID
	TemplateID ID
	UserID     ID
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

// IsExpired проверяет, истекла ли сессия тестирования
func (s *TestSession) IsExpired() bool {
	// Логика: если  ExpiredAt в прошлом, то сессия истекла
	return s.ExpiredAt.Before(time.Now())
}
