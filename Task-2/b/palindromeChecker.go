package b

import (
	"unicode"
)

func IsPalindrome(text string) bool {
	var filtered []rune

	for _,char := range text {
		if char >= 'A' && char <= 'Z' {
			char += 'a' - 'A'
		}

		if unicode.IsLetter(char) || unicode.IsDigit(char) {
			filtered = append(filtered, char)
		}
	}

	n := len(filtered)
	for i := 0; i < n/2; i++ {
		if filtered[i] != filtered[n-i-1] {
			return false
		}
	}
	return true
}