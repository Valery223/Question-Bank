package domain

import "testing"

func TestDifficulty(t *testing.T) {
	tests := []struct {
		name       string
		difficulty Difficulty
		valid      bool
	}{
		{"Valid difficulty 1", 1, true},
		{"Valid difficulty 3", 3, true},
		{"Valid difficulty 5", 5, true},
		{"Invalid difficulty 0", 0, false},
		{"Invalid difficulty 6", 6, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.difficulty.IsValid(); got != tt.valid {
				t.Errorf("IsValid() = %v, want %v", got, tt.valid)
			}
		})
	}
}
