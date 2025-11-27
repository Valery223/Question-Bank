package domain

// Option - Вариант ответа
type Option struct {
	ID int
	// QuestionID int - В Go внутри домена это поле часто не нужно,
	// так как Option всегда "живет" внутри Question. Но для SQL удобно иметь.
	Text      string
	IsCorrect bool
}

// Question - Агрегат (Корень)
type Question struct {
	ID         int
	Role       RoleQuestionnaire
	Topic      string
	Type       QuestionType
	Difficulty Difficulty
	Text       string

	// Композиция: Вопрос владеет своими опциями
	Options []Option
}

// Validate проверяет инварианты сущности
func (q *Question) Validate() error {
	if q.Text == "" {
		return ErrEmptyText
	}

	if !q.Difficulty.IsValid() {
		return ErrInvalidDifficulty
	}

	if !q.Type.IsValid() {
		return ErrInvalidType
	}

	switch q.Type {
	case TypeSingleChoice, TypeMultiChoice:
		if len(q.Options) < 2 {
			return ErrNoOptions
		}
	case TypeText:
		// Текстовые вопросы не требуют опций
	default:
		return ErrInvalidType
	}

	for _, option := range q.Options {
		if err := option.Validate(); err != nil {
			return err
		}
	}

	return nil
}

func (o *Option) Validate() error {
	if o.Text == "" {
		return ErrEmptyText
	}
	return nil
}
