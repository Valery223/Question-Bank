package domain

type TestTemplate struct {
	ID      string
	Name    string
	Role    RoleQuestionnaire
	Purpose TemplatePurpose

	Questions []Question
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
