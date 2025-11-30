package v1

import "github.com/Valery223/Question-Bank/internal/domain"

// --- Requests (Входящие данные) ---

type OptionRequest struct {
	Text     string `json:"text" binding:"required"`
	IsAnswer bool   `json:"is_answer"` // false по умолчанию
}

type CreateQuestionRequest struct {
	Role       string          `json:"role" binding:"required"`
	Topic      string          `json:"topic" binding:"required"`
	Type       string          `json:"type" binding:"required"`
	Difficulty int             `json:"difficulty" binding:"required"`
	Text       string          `json:"text" binding:"required"`
	Options    []OptionRequest `json:"options" binding:"required,min=2"`
}

type UpdateQuestionRequest struct {
	Role       string          `json:"role" `
	Topic      string          `json:"topic" `
	Type       string          `json:"type" `
	Difficulty int             `json:"difficulty" `
	Text       string          `json:"text" `
	Options    []OptionRequest `json:"options" binding:"min=2"`
}

type CreateTemplateRequest struct {
	Name        string   `json:"name" binding:"required"`
	Role        string   `json:"role" binding:"required"`
	Purpose     string   `json:"purpose" binding:"required"`
	QuestionIDs []string `json:"question_ids" binding:"required,min=1"`
}

type UpdateTemplateRequest struct {
	Name        string   `json:"name" `
	Role        string   `json:"role" `
	Purpose     string   `json:"purpose" `
	QuestionIDs []string `json:"question_ids" binding:"min=1"`
}

type CreateTestSessionRequest struct {
	TemplateID string `json:"template_id" binding:"required"`
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

type TestSessionResponse struct {
	ID        string             `json:"session_id"`
	Questions []QuestionResponse `json:"questions"`
}

// Ответ для шаблона теста
//
// Включает список ID вопросов
// Если нужны будут полные вопросы, то можно всделать отдельный эндпоинт для получения вопросов по шаблону
// Напри
type TemplateResponse struct {
	ID          string   `json:"id"`
	Name        string   `json:"name"`
	Role        string   `json:"role"`
	Purpose     string   `json:"purpose"`
	QuestionIDs []string `json:"question_ids"` // Список ID вопросов
}

type TemplateDetailsResponse struct {
	ID        string             `json:"id"`
	Name      string             `json:"name"`
	Role      string             `json:"role"`
	Purpose   string             `json:"purpose"`
	Questions []QuestionResponse `json:"questions"` // Полные вопросы
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

func (r *UpdateQuestionRequest) ToDomain() *domain.Question {
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

func (t *CreateTemplateRequest) ToDomain() *domain.TestTemplate {
	td := &domain.TestTemplate{
		Name:        t.Name,
		Role:        domain.RoleQuestionnaire(t.Role),
		Purpose:     domain.TemplatePurpose(t.Purpose),
		QuestionIDs: make([]domain.ID, len(t.QuestionIDs)),
	}

	for i, id := range t.QuestionIDs {
		td.QuestionIDs[i] = domain.ID(id)
	}

	return td
}

func (r *CreateTestSessionRequest) ToDomain() *domain.TestSession {
	ts := &domain.TestSession{
		TemplateID: domain.ID(r.TemplateID),
	}

	return ts
}

func (t *UpdateTemplateRequest) ToDomain() *domain.TestTemplate {
	td := &domain.TestTemplate{
		Name:        t.Name,
		Role:        domain.RoleQuestionnaire(t.Role),
		Purpose:     domain.TemplatePurpose(t.Purpose),
		QuestionIDs: make([]domain.ID, len(t.QuestionIDs)),
	}

	for i, id := range t.QuestionIDs {
		td.QuestionIDs[i] = domain.ID(id)
	}

	return td
}

// ---
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

func TemplateToResponse(t *domain.TestTemplate) *TemplateResponse {
	return &TemplateResponse{
		ID:      string(t.ID),
		Name:    t.Name,
		Role:    string(t.Role),
		Purpose: string(t.Purpose),
		QuestionIDs: func() []string {
			ids := make([]string, len(t.QuestionIDs))
			for i, id := range t.QuestionIDs {
				ids[i] = string(id)
			}
			return ids
		}(),
	}
}

func TemplateDetailsToResponse(t *domain.TestTemplate, questions []domain.Question) *TemplateDetailsResponse {
	resp := &TemplateDetailsResponse{
		ID:        string(t.ID),
		Name:      t.Name,
		Role:      string(t.Role),
		Purpose:   string(t.Purpose),
		Questions: make([]QuestionResponse, len(questions)),
	}

	for i, q := range questions {
		resp.Questions[i] = *QuestionToResponse(&q)
	}

	return resp
}

func TestSessionToResponse(s *domain.TestSession) *TestSessionResponse {
	resp := &TestSessionResponse{
		ID:        string(s.ID),
		Questions: make([]QuestionResponse, len(s.Questions)),
	}

	for i, q := range s.Questions {
		resp.Questions[i] = *QuestionToResponse(&q)
	}

	return resp
}
