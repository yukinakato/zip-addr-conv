package main

import (
	"reflect"
	"testing"
)

func TestSubstr(t *testing.T) {
	type args struct {
		s string
		n int
	}
	cases := []struct {
		arg      args
		expected string
	}{
		{args{"hello", 0}, "hello"},
		{args{"hello", 2}, "he"},
		{args{"hello", -1}, "hello"},
		{args{"hello", 100}, "hello"},
		{args{"", 100}, ""},
		{args{"", 0}, ""},
	}

	for _, tt := range cases {
		actual := substr(tt.arg.s, tt.arg.n)
		if actual != tt.expected {
			t.Errorf("substr(%s, %d) ... expected %s, actual %s", tt.arg.s, tt.arg.n, tt.expected, actual)
		}
	}
}

func TestBuildKeywordArray(t *testing.T) {
	cases := []struct {
		in       string
		expected []string
	}{
		{
			"hello",
			[]string{"hello"},
		},
		{
			"he llo",
			[]string{"he", "llo"},
		},
		{
			"he　l-l　 - 　　o-",
			[]string{"he", "ll", "o"},
		},
		{
			"　　 　hel-lo 1２３-45-６７ ８90  　　 ",
			[]string{"hello", "1234567", "890"},
		},
		{
			"　　 　１２３４５ ６７８９０  　　 ",
			[]string{"12345", "67890"},
		},
	}

	for _, tt := range cases {
		actual := buildKeywordArray(tt.in)
		if !reflect.DeepEqual(tt.expected, actual) {
			t.Errorf("buildKeywordArray(%s) ... expected %s, actual %s", tt.in, tt.expected, actual)
		}
	}
}
