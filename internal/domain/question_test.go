package domain

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQuestion(t *testing.T) {
	tests := []struct {
		name      string
		question  Question
		wantError error
	}{
		{
			name: "Valid single choice question",
			question: Question{
				Text:       "What is 2 + 2?",
				Type:       TypeSingleChoice,
				Difficulty: 1,
				Options: []Option{
					{Text: "3", IsCorrect: false},
					{Text: "4", IsCorrect: true},
				},
			},
			wantError: nil,
		},
		{
			name: "Invalid question with empty text",
			question: Question{
				Text:       "",
				Type:       TypeText,
				Difficulty: 2,
			},
			wantError: ErrEmptyText,
		},
		{
			name: "Invalid question with insufficient options",
			question: Question{
				Text:       "Choose the correct option",
				Type:       TypeMultiChoice,
				Difficulty: 3,
				Options: []Option{
					{Text: "Option 1", IsCorrect: false},
				},
			},
			wantError: ErrNoOptions,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.question.Validate()
			// if err != tt.wantError {
			// 	t.Errorf("Validate() = %v, want %v", err, tt.wantError)
			// }
			if tt.wantError == nil {
				assert.NoError(t, err, "expected no error, but got one")
			} else {
				assert.EqualError(t, err, tt.wantError.Error(), "expected error does not match actual error")
			}
		})
	}
}
