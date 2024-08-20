package concat_test

import (
	concat "concat/internal"
	"testing"
)

func TestConcat(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		sep      string
		split    string
		expected string
	}{
		{
			name:     "ConcatWithSpace",
			input:    "Hello World",
			sep:      " ",
			split:    " ",
			expected: "hello world",
		},
		{
			name:     "ConcatWithComma",
			input:    "apple,banana,orange",
			sep:      "-",
			split:    ",",
			expected: "apple-banana-orange",
		},
		{
			name:     "ConcatWithEmptyString",
			input:    "",
			sep:      "-",
			split:    ",",
			expected: "",
		},
		{
			name:     "ConcatWithSingleChar",
			input:    "a",
			sep:      "-",
			split:    ",",
			expected: "a",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := concat.Concat(tc.input, tc.sep, tc.split)
			if result != tc.expected {
				t.Errorf("Concat(%q, %q, %q) = %q; expected %q", tc.input, tc.sep, tc.split, result, tc.expected)
			}
		})
	}
}
