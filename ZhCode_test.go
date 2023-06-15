package main

import (
	"testing"
)

// golang.org\x\text\encoding\simplifiedchinese\all_test.go
func TestGbkEncode1(t *testing.T) {
	s := `花间一壶酒，独酌无相亲。`
	ss := GbkEncode(s)
	res := len(ss)
	expected := 24

	if res != expected {
		t.Errorf("Result: %v, want: %v", res, expected)
	}
}

func TestGbkEnDecode1(t *testing.T) {
	s := `花间一壶酒，独酌无相亲。`
	ss := GbkEncode(s)
	res := GbkDecode(ss)
	expected := s

	if res != expected {
		t.Errorf("Result: %v, want: %v", res, expected)
	}
}
