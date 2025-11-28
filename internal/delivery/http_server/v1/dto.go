package v1

import "github.com/Valery223/Question-Bank/internal/domain"

// --- Requests (Входящие данные) ---

type CreateQuestionRequest struct {
	Role       string          `json:"role" binding:"required"`
	Topic      string          `json:"topic" binding:"required"`
	Type       string          `json:"type" binding:"required"`
	Difficulty int             `json:"difficulty" binding:"required"`
	Text       string          `json:"text" binding:"required"`
	Options    []OptionRequest `json:"options"`
}

type OptionRequest struct {
	Text     string `json:"text" binding:"required"`
	IsAnswer bool   `json:"is_answer" binding:"required"`
}

// --- Responses (Исходящие данные) ---

type QuestionResponse struct {
	ID         string           `json:"id"`
	Role       string           `json:"role"`
	Topic      string           `json:"topic"`
	Type       string           `json:"type"`
	Difficulty int              `json:"difficulty"`
	Text       string           `json:"text"`
	Options    []OptionResponse `json:"options,omitempty"`
}

type OptionResponse struct {
	ID       string `json:"id"`
	Text     string `json:"text"`
	IsAnswer bool   `json:"is_answer"`
}

// --- Mappers (Преобразователи DTO <-> Domain) ---
// Хендлер не должен писать маппинг на 50 строк. Вынесем это в методы DTO.

func (r *CreateQuestionRequest) ToDomain() *domain.Question {
	q := &domain.Question{
		Role:       domain.RoleQuestionnaire(r.Role),
		Topic:      r.Topic,
		Type:       domain.QuestionType(r.Type),
		Difficulty: domain.Difficulty(r.Difficulty),
		Text:       r.Text,
		Options:    make([]domain.Option, len(r.Options)),
	}

	for i, opt := range r.Options {
		q.Options[i] = domain.Option{
			Text:      opt.Text,
			IsCorrect: opt.IsAnswer,
		}
	}

	return q
}

// Из домена в DTO ответа
func QuestionToResponse(q *domain.Question) *QuestionResponse {
	resp := &QuestionResponse{
		ID:         string(q.ID),
		Role:       string(q.Role),
		Topic:      q.Topic,
		Type:       string(q.Type),
		Difficulty: int(q.Difficulty),
		Text:       q.Text,
		Options:    make([]OptionResponse, len(q.Options)),
	}

	for i, opt := range q.Options {
		resp.Options[i] = OptionResponse{
			ID:       string(opt.ID),
			Text:     opt.Text,
			IsAnswer: opt.IsCorrect,
		}
	}

	return resp
}
