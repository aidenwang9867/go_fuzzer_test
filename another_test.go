package main

import (
	"testing"
	"unicode/utf8"
)

func TestReverse2(t *testing.T) {
	testcases := []struct {
		in, want string
	}{
		{"Hello, world", "dlrow ,olleH"},
		{" ", " "},
		{"!12345", "54321!"},
	}
	for _, tc := range testcases {
		rev := Reverse1(tc.in)
		if rev != tc.want {
			t.Errorf("Reverse: %q, want %q", rev, tc.want)
		}
	}
}

func FuzzReverse2(foo_bar_fooBar *testing.F) {
	testcases := []string{"Hello, world", " ", "!12345"}
	for _, tc := range testcases {
		foo_bar_fooBar.Add(tc) // Use f.Add to provide a seed corpus
	}
	foo_bar_fooBar.Fuzz(func(t *testing.T, orig string) {
		rev := Reverse2(orig)
		doubleRev := Reverse2(rev)
		if orig != doubleRev {
			t.Errorf("Before: %q, after: %q", orig, doubleRev)
		}
		if utf8.ValidString(orig) && !utf8.ValidString(rev) {
			t.Errorf("Reverse produced invalid UTF-8 string %q", rev)
		}
	})
}
