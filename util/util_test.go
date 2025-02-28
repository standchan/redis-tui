package util

import "testing"

func TestParseTime(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
		wantErr  bool
	}{
		{
			name:     "seconds only",
			input:    "45",
			expected: "45s",
			wantErr:  false,
		},
		{
			name:     "minutes and seconds",
			input:    "125",
			expected: "2m5s",
			wantErr:  false,
		},
		{
			name:     "hours and minutes",
			input:    "3720",
			expected: "1h2m",
			wantErr:  false,
		},
		{
			name:     "days and hours",
			input:    "93600",
			expected: "1d2h",
			wantErr:  false,
		},
		{
			name:     "invalid input",
			input:    "abc",
			expected: "",
			wantErr:  true,
		},
		{
			name:     "zero seconds",
			input:    "0",
			expected: "0s",
			wantErr:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := FormatUptime(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseTime() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.expected {
				t.Errorf("ParseTime() = %v, want %v", got, tt.expected)
			}
		})
	}
}
