package wcwidth

import (
	"testing"
)

func TestWcwidth(t *testing.T) {
	var tests = []struct {
		in  rune
		out int
	}{
		{'a', 1}, // U+0061
		{'ア', 2}, // U+30A2, fullwidth Katakana
		{'ｱ', 1}, // U+FF71, halfwidth Katakana
		{'★', 2}, // U+2605, ambiguous
		{'¢', 2}, // U+00A2, legacy japanese
	}

	for _, tt := range tests {
		if out := Wcwidth(tt.in); out != tt.out {
			t.Errorf("expected width of %q: %v, actual: %v", tt.in, tt.out, out)
		}
	}
}

func TestWcwidthUcs(t *testing.T) {
	var tests = []struct {
		in  rune
		out int
	}{
		{'a', 1}, // U+0061
		{'ア', 2}, // U+30A2, fullwidth Katakana
		{'ｱ', 1}, // U+FF71, halfwidth Katakana
		{'★', 1}, // U+2605, ambiguous
		{'¢', 1}, // U+00A2, legacy japanese
	}

	for _, tt := range tests {
		if out := WcwidthUcs(tt.in); out != tt.out {
			t.Errorf("expected width of %q: %v, actual: %v", tt.in, tt.out, out)
		}
	}
}

func TestWcswidth(t *testing.T) {
	var tests = []struct {
		in  string
		out int
	}{
		{"aあアｱ☆¢", 10},
	}

	for _, tt := range tests {
		if out := Wcswidth(tt.in); out != tt.out {
			t.Errorf("expected width of %q: %v, actual: %v", tt.in, tt.out, out)
		}
	}
}

func TestWcswidthUcs(t *testing.T) {
	var tests = []struct {
		in  string
		out int
	}{
		{"aあアｱ☆¢", 8},
	}

	for _, tt := range tests {
		if out := WcswidthUcs(tt.in); out != tt.out {
			t.Errorf("expected width of %q: %v, actual: %v", tt.in, tt.out, out)
		}
	}
}
