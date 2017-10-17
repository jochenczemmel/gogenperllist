package main

import (
	"testing"
)

func TestExtractPrefix(t *testing.T) {
	candidates := []struct {
		in       string
		expected string
	}{
		{"github.com/jochenczemmel/mypackage", "mypackage."},
		{"github.com/jochenczemmel/mypackage.v2", "mypackage."},
		{"github.com/jochenczemmel/my-package", "package."},
		{"github.com/jochenczemmel/my-package.v2", "package."},
		{"mypackage", "mypackage."},
	}

	for _, c := range candidates {
		got := extractPrefix(c.in)
		if got != c.expected {
			t.Errorf("NOTOK: got %s, expected %s\n", got, c.expected)
		}
	}
}
