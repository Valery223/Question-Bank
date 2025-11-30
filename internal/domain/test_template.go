package domain

// TestTemplate представляет шаблон теста с набором вопросов
type TestTemplate struct {
	ID      ID
	Name    string
	Role    RoleQuestionnaire
	Purpose TemplatePurpose

	// QuestionIDs содержит список ID вопросов, входящих в шаблон
	//
	// Порядок вопросов важен
	QuestionIDs []ID // Список ID вопросов
}

// Validate проверяет инварианты сущности
func (tt *TestTemplate) Validate() error {
	if tt.Name == "" {
		return ErrEmptyText
	}

	if !tt.Role.IsValid() {
		return ErrInvalidType
	}

	if !tt.Purpose.IsValid() {
		return ErrInvalidType
	}

	return nil
}
