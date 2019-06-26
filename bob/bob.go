// Package bob provides ...

/* I need help and have some question. Check it please.
1. What should I write in package comment? what is it for or how does it work?
2. I specifically did not use regexp to complete this task,
	because I wanted to understand and remember	functions of strings, unicode packages.
	So I can`t pass test for 100%. Maybe you tell me where I was wrong?! Thanks a lot.
*/

package bob

import (
	"strings"
	"unicode"
)

const suffixQ = "?"
const suffixY = "!"

// Hey answers to Bob depending on functions below
func Hey(s string) string {
	s = strings.TrimSpace(strings.Replace(s, " ", "", -1))

	if isQuestion(s) {
		return "Sure."
	} else if isYelling(s) {
		return "Whoa, chill out!"
	} else if isYellingQuestion(s) {
		return "Calm down, I know what I'm doing!"
	} else if isSilence(s) {
		return "Fine. Be that way!"
	}
	return "Whatever."
}

func isQuestion(s string) bool {
	return strings.HasSuffix(s, suffixQ) && s != strings.ToUpper(s)
}

func isYelling(s string) bool {
	return strings.HasSuffix(s, suffixY) && s == strings.ToUpper(s) || isLetter(s)
}

func isSilence(s string) bool {
	return len(s) == 0
}

func isYellingQuestion(s string) bool {
	return strings.HasSuffix(s, suffixQ) && s == strings.ToUpper(s) && !isDigit(s)
}

func isLetter(s string) bool {
	for i := 0; i < len(s); i++ {
		if !unicode.IsLetter(rune(s[i])) {
			return false
		}
	}
	return true
}

func isDigit(s string) bool {
	for _, r := range s {
		if !unicode.IsDigit(r) {
			return false
		}
	}
	return true
}
