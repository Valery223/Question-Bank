package domain

import (
	"testing"
	"time"
)

func TestTestSession(t *testing.T) {
	tests := []struct {
		name        string
		testSession TestSession
		wantExpired bool
	}{
		{
			name: "Expired session",
			testSession: TestSession{
				ExpiredAt: time.Now().Add(-1 * time.Hour),
			},
			wantExpired: true,
		},
		{
			name: "Active session",
			testSession: TestSession{
				ExpiredAt: time.Now().Add(1 * time.Hour),
			},
			wantExpired: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.testSession.IsExpired(); got != tt.wantExpired {
				t.Errorf("IsExpired() = %v, want %v", got, tt.wantExpired)
			}
		})
	}
}
