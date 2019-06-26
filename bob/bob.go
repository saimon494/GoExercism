// Package bob
package bob

import (
	"strings"
	"unicode"
)

const suffixQ = "?"
const suffixY = "!"

// Hey answers to Bob depending on 3 functions below
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

// s == strings.ToLower(s)
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

// func Hey(remark string) string {
//   remark = strings.TrimSpace(remark)
//
//   lowerRe, _ := regexp.Compile("[a-z]")
//   anyLettersRe, _ := regexp.Compile("[a-zA-Z]")
//
//   isSilence := len(remark) == 0
//   isQuestion := strings.HasSuffix(remark, "?")
//   isExclamation := !lowerRe.MatchString(remark)
//   anyLetters := anyLettersRe.MatchString(remark)
//
//   if isSilence {
//     return "Fine. Be that way!"
//   } else if !anyLetters && isQuestion {
//     return "Sure."
//   } else if !anyLetters {
//     return "Whatever."
//   } else if isQuestion && isExclamation {
//     return "Calm down, I know what I'm doing!"
//   } else if isQuestion {
//     return "Sure."
//   } else if isExclamation {
//     return "Whoa, chill out!"
//   } else {
//     return "Whatever."
//   }
// }
