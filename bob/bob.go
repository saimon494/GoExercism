// Package bob
package bob

import (
	"strings"
	"unicode"
)

const suffix = "?"

// Hey answers to Bob depending on 3 functions below
func Hey(s string) string {
	s = strings.TrimSpace(s)

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
	// suffix := "?"
	return strings.HasSuffix(s, suffix) && s != strings.ToUpper(s)
}

// s == strings.ToLower(s)
func isYelling(s string) bool {
	return !strings.HasSuffix(s, suffix) && s == strings.ToUpper(s) && isLetter(s)
}

func isSilence(s string) bool {
	return len(s) == 0
}

func isYellingQuestion(s string) bool {
	return s == strings.ToUpper(s) && strings.HasSuffix(s, suffix)
}

func isLetter(s string) bool {
	for _, r := range s {
		if !unicode.IsLetter(r) {
			return false
		}
	}
	return true
}

// && !unicode.IsPunct(r) && !unicode.IsGraphic(r)

// func noLetter(s string) bool {
// 	for _, r := range s {
// 		if unicode.IsSpace(r) && unicode.IsPunct(r) {
// 			return true
// 		}
// 	}
// 	return false
// }

// func isGraphic(s string) bool {
// 	for _, r := range s {
// 		if !unicode.IsGraphic(r) {
// 			return false
// 		}
// 	}
// 	return true
// }
