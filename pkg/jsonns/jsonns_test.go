package jsonns

import (
	"regexp"
	"testing"
)

func TestNSHasKeyArrayIndex(t *testing.T) {
	tests := []struct {
		input string
		want  bool
	}{
		{"names[0]", true},
		{"names[a]", false},
		{"names", false},
	}

	for _, tt := range tests {
		got := NSHasKeyArrayIndex(tt.input)
		if got != tt.want {
			t.Errorf("NSHasKeyArrayIndex(%q) = %v, want %v", tt.input, got, tt.want)
		}
	}
}

func TestNSGetKeyArrayIndex(t *testing.T) {
	tests := []struct {
		input string
		want  *string
	}{
		{"names[0]", strPtr("0")},
		{"names[name]", nil},
		{"names", nil},
	}

	for _, tt := range tests {
		got := NSGetKeyArrayIndex(tt.input)
		if (got == nil && tt.want != nil) || (got != nil && tt.want == nil) || (got != nil && *got != *tt.want) {
			t.Errorf("NSGetKeyArrayIndex(%q) = %v, want %v", tt.input, got, tt.want)
		}
	}
}

func TestNSGetKeyArrKey(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{"names[0]", "names"},
		{"names[name]", "names[name]"},
	}

	for _, tt := range tests {
		got := NSGetKeyArrKey(tt.input)
		if got != tt.want {
			t.Errorf("NSGetKeyArrKey(%q) = %v, want %v", tt.input, got, tt.want)
		}
	}
}

func TestNSKeyarrObjify(t *testing.T) {
	tests := []struct {
		input string
		want  KeyArrObj
	}{
		{"names[0]", KeyArrObj{Key: "names", Index: strPtr("0")}},
		{"names[name]", KeyArrObj{Key: "names[name]", Index: nil}},
	}

	for _, tt := range tests {
		got := NSKeyarrObjify(tt.input)
		if got.Key != tt.want.Key || !equalPtrString(got.Index, tt.want.Index) {
			t.Errorf("NSKeyarrObjify(%q) = %v, want %v", tt.input, got, tt.want)
		}
	}
}

func TestNSPureName(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{"[0]", "0"},
		{"[name]", "name"},
	}

	for _, tt := range tests {
		got := NSPureName(tt.input)
		if got != tt.want {
			t.Errorf("NSPureName(%q) = %v, want %v", tt.input, got, tt.want)
		}
	}
}

func TestNSGetMatch(t *testing.T) {
	tests := []struct {
		input  string
		reg    *regexp.Regexp
		preset string
		want   []string
	}{
		{"names[0]", nil, "", []string{"[0]"}},
		{"names[0][1]", nil, "", []string{"[0]", "[1]"}},
		{"names", nil, "", nil},
		{"names[zero]", nil, "", []string{"[zero]"}},
		{"names[zero]", regexp.MustCompile(`\[\d+\]`), "", nil},
	}

	for _, tt := range tests {
		got := NSGetMatch(tt.input, tt.reg, tt.preset)
		if !equalStringSlices(got, tt.want) {
			t.Errorf("NSGetMatch(%q, %v, %q) = %v, want %v", tt.input, tt.reg, tt.preset, got, tt.want)
		}
	}
}

func TestNSStd(t *testing.T) {
	tests := []struct {
		input string
		sep   string
		reg   *regexp.Regexp
		want  []string
	}{
		{"names[zero]", ".", nil, []string{"names", "[zero]"}},
		{"a.b.c", ".", nil, []string{"a", "b", "c"}},
	}

	for _, tt := range tests {
		got := NSStd(tt.input, tt.sep, tt.reg)
		if !equalStringSlices(got, tt.want) {
			t.Errorf("NSStd(%q, %q, %v) = %v, want %v", tt.input, tt.sep, tt.reg, got, tt.want)
		}
	}
}

// Helper functions
func strPtr(s string) *string {
	return &s
}

func equalPtrString(a, b *string) bool {
	if a == nil && b == nil {
		return true
	}
	if a == nil || b == nil {
		return false
	}
	return *a == *b
}

func equalStringSlices(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
} 