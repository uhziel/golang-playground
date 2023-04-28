package word

import "unicode"

func IsPalindrome(s string) bool {
	runes := make([]rune, 0, len(s))
	for _, r := range s {
		if unicode.IsLetter(r) {
			runes = append(runes, unicode.ToLower(r))
		}
	}

	n := len(runes)
	for i := 0; i < n/2; i++ {
		if runes[i] != runes[n-i-1] {
			return false
		}
	}

	return true
}
