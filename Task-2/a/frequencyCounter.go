package a

import (
	"unicode"
)

func FrequencyCount(text string) map[string]int {
	freq := make(map[string]int)

	var lowerText string
    for _, ch := range text {
	    if ch >= 'A' && ch <= 'Z' {
		    ch += ('a' - 'A')
	    }
	    lowerText += string(ch)
    }

	word := ""
    for _, ch := range lowerText {
	    if unicode.IsLetter(ch) || unicode.IsDigit(ch) {
		    word += string(ch)
	    } else if word != "" {
		    freq[word]++
		    word = ""
	    }
    }
    if word != "" {
	    freq[word]++
    }

	return freq
}
