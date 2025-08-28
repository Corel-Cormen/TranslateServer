package SentenceFormatterImpl

import (
	"testing"
)

func TestPrepareInput(t *testing.T) {
	formatter := &MarianSentenceFormatter{}

	tests := []struct {
		input    string
		expected string
	}{
		{
			input:    "My mother cooks soup for dinner. The soup is hot. Then I go to bed",
			expected: "My mother cooks soup for dinner . The soup is hot . Then I go to bed .",
		},
		{
			input:    "Hello, world!",
			expected: "Hello , world !",
		},
		{
			input:    "Test without dot",
			expected: "Test without dot .",
		},
	}

	for _, tt := range tests {
		out := formatter.PrepareInput(tt.input)
		if out != tt.expected {
			t.Errorf("PrepareInput(%q) = %q, want %q", tt.input, out, tt.expected)
		}
	}
}

func TestCleanOutput(t *testing.T) {
	formatter := &MarianSentenceFormatter{}

	tests := []struct {
		input    string
		expected string
	}{
		{
			input:    "▁Moja ▁matka ▁lubi ▁koty .",
			expected: "Moja matka lubi koty.",
		},
		{
			input:    "To ▁jest ▁test ,przykład !",
			expected: "To jest test, przykład!",
		},
		{
			input:    "Lu bi ę ▁chodzić ▁do ▁szkoły .",
			expected: "Lubię chodzić do szkoły.",
		},
	}

	for _, tt := range tests {
		out := formatter.CleanOutput(tt.input)
		if out != tt.expected {
			t.Errorf("CleanOutput(%q) = %q, want %q", tt.input, out, tt.expected)
		}
	}
}
