package main

import (
	"strings"
	"testing"
)

func TestEcho(t *testing.T) {
	tests := []struct {
		newline bool
		sep     string
		args    []string
		want    string
	}{
		{true, " ", []string{}, "\n"},
		{false, " ", []string{}, ""},
		{true, " ", []string{"a", "b", "c"}, "a b c\n"},
		{true, "\t", []string{"a", "b", "c"}, "a\tb\tc\n"},
		{true, ",", []string{"a", "b", "c"}, "a,b,c\n"},
	}

	for _, test := range tests {
		output = new(strings.Builder)
		if err := echo(test.newline, test.sep, test.args); err != nil {
			t.Errorf("%s", err)
			continue
		}

		ans := output.(*strings.Builder).String()
		if ans != test.want {
			t.Errorf("echo(%v, %q, %q) = %q, want %q", test.newline, test.sep, test.args, ans, test.want)
		}
	}
}
