package word

import "testing"

func TestPalindrome(t *testing.T) {
	cases := []struct {
		input    string
		expected bool
	}{
		{"abc", false},
		{"aba", true},
		{"", true},
		{"foobar", false},
		{"refer", true},
	}

	for _, c := range cases {
		if IsPalindrome(c.input) != c.expected {
			t.Errorf("isPalindrome(%q) expected:%v real:%v", c.input, c.expected, IsPalindrome(c.input))
		}
	}
}
