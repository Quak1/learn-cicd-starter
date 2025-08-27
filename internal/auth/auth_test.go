package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	tests := map[string]struct {
		key         string
		value       string
		output      string
		expectError bool
	}{
		"good Authorization header": {
			key:         "Authorization",
			value:       "ApiKey testApiKey",
			output:      "testApiKey",
			expectError: false,
		},
		"no Authorization": {
			output:      "",
			expectError: true,
		},
		"short Authorization": {
			key:         "Authorization",
			value:       "ApiKey",
			output:      "",
			expectError: true,
		},
		"typo Authorization": {
			key:         "Authorization",
			value:       "ApiKe testApiKey",
			output:      "",
			expectError: true,
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			header := http.Header{}
			header.Add(tc.key, tc.value)

			output, err := GetAPIKey(header)
			if err != nil {
				if !tc.expectError {
					t.Errorf("Unexpected error")
				}
				return
			}

			if output != tc.output {
				t.Errorf("Expected: %s, but got: %s", tc.output, output)
			}
		})
	}
}
