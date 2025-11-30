package v1

import (
	"bytes"
	"io"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Valery223/Question-Bank/internal/mocks"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestHandler_CreateQuestion(t *testing.T) {
	// Тишина в логах
	logger := slog.New(slog.NewTextHandler(io.Discard, nil))

	validRequestBody := `{
		"role": "backend_junior",
		"topic": "math",
		"type": "multiple_choice",
		"difficulty": 1,
		"text": "What is 2 + 2?",
		"options": [
			{"text": "3", "is_answer": false},
			{"text": "4", "is_answer": true}
		]
	}`

	// Невалидный запрос
	invalidRequestBody := `{
		"role": "backend_junior",
		"topic": "math",
		"type": "multiple_choice",
		"difficulty": 1,
		"text1": "sasas",
		"options": [
			{"text": "3", "is_answer": false},
			{"text": "4", "is_answer": true}
		]
	}`

	// Пустой мок TemplateUseCase и SessionUseCase

	tests := []struct {
		name           string
		requestBody    string
		setupMock      func(m *mocks.MockQuestionUseCase) // Предположим, мы сделали интерфейс для UC
		expectedStatus int
		expectedBody   string
	}{
		{
			name:        "Successful question creation",
			requestBody: validRequestBody,
			setupMock: func(m *mocks.MockQuestionUseCase) {
				m.On("CreateQuestion", mock.Anything, mock.AnythingOfType("*domain.Question")).Return(nil).Once()
			},
			expectedStatus: http.StatusCreated,
			expectedBody:   `"status":"created"`,
		},
		{
			name:        "Invalid request body",
			requestBody: invalidRequestBody,
			setupMock: func(m *mocks.MockQuestionUseCase) {
				// Мок не должен вызываться
			},
			expectedStatus: http.StatusBadRequest,
			expectedBody:   "invalid request body",
		},
		{
			name:        "UseCase returns error",
			requestBody: validRequestBody,
			setupMock: func(m *mocks.MockQuestionUseCase) {
				m.On("CreateQuestion", mock.Anything, mock.AnythingOfType("*domain.Question")).Return(assert.AnError).Once()
			},
			expectedStatus: http.StatusInternalServerError,
			expectedBody:   assert.AnError.Error(),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Создаем мок для QuestionUseCase
			mockQuestionUC := mocks.NewMockQuestionUseCase(t)

			tt.setupMock(mockQuestionUC) // Остальные моки не используются в этом тесте

			// Создаем хендлер с моками
			h := NewHandler(mockQuestionUC, nil, nil, logger)

			// Подготовка Gin
			gin.SetMode(gin.TestMode)
			r := gin.New()
			r.POST("/questions", h.CreateQuestion)

			// Создаем HTTP Запрос
			req := httptest.NewRequest(http.MethodPost, "/questions", bytes.NewBufferString(tt.requestBody))
			req.Header.Set("Content-Type", "application/json")

			// Recorder
			w := httptest.NewRecorder()

			// Пуск
			r.ServeHTTP(w, req)

			//  Проверки
			assert.Equal(t, tt.expectedStatus, w.Code)
			if tt.expectedBody != "" {
				// Проверяем, что тело ответа содержит ожидаемую строку
				assert.Contains(t, w.Body.String(), tt.expectedBody)
			}

			mockQuestionUC.AssertExpectations(t)
		})
	}
}
