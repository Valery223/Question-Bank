package domain

import "testing"

// Название крутое
func TestTestTemplate(t *testing.T) {
	tests := []struct {
		name         string
		testTemplate TestTemplate
		wantError    error
	}{
		{
			name: "Valid TestTemplate",
			testTemplate: TestTemplate{
				Name:    "Sample Test",
				Role:    RoleBackendJunior,
				Purpose: PurposeAssessment,
			},
		},
		{
			name: "Invalid TestTemplate with empty name",
			testTemplate: TestTemplate{
				Name:    "",
				Role:    RoleBackendJunior,
				Purpose: PurposeAssessment,
			},
			wantError: ErrEmptyText,
		},
		{
			name: "Invalid TestTemplate with invalid role",
			testTemplate: TestTemplate{
				Name:    "Sample Test",
				Role:    "invalid_role",
				Purpose: PurposeAssessment,
			},
			wantError: ErrInvalidType,
		},
		{
			name: "Invalid TestTemplate with invalid purpose",
			testTemplate: TestTemplate{
				Name:    "Sample Test",
				Role:    RoleBackendJunior,
				Purpose: "invalid_purpose",
			},
			wantError: ErrInvalidType,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.testTemplate.Validate()
			if err != tt.wantError {
				t.Errorf("Validate() = %v, want %v", err, tt.wantError)
			}
		})
	}
}
