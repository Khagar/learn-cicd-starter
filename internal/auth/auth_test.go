package auth

import (
	"errors"
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	tests := []struct {
		name      string
		headers   http.Header
		expected  string
		expectErr error
	}{
		{
			name:      "No Authorization header",
			headers:   http.Header{},
			expected:  "",
			expectErr: ErrNoAuthHeaderIncluded,
		},
		{
			name:      "Valid Authorization header",
			headers:   http.Header{"Authorization": []string{"ApiKey my-api-key"}},
			expected:  "my-api-key",
			expectErr: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := GetAPIKey(tt.headers)
			if result != tt.expected || !errors.Is(err, tt.expectErr) {
				t.Errorf("GetAPIKey() = %v, %v; want %v, %v", result, err, tt.expected, tt.expectErr)
			}
		})
	}
}
