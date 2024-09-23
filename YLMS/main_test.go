package printer

import "testing"

func TestDeleteVowels(t *testing.T) {
	tests := []struct {
		input string
		expected string
	}{
		{"hello", "hll"},
        {"world", "wrld"},
        {"aeiou", ""},
        {"Go is fun", "G s fn"},
        {"ABCDEFGHIJKLMNOPQRSTUVWXYZ", "BCDFGHJKLMNPQRSTVWXYZ"},
        {"", ""},  
        {"xyz", "xyz"},
	}

	for _, tt := range tests {
		got := DeleteVowels(tt.input)
		if got != tt.expected {
			t.Fatalf("DeleteVowels(%q) = %q; want %q", tt.input, got, tt.expected)
		}
	}
}
