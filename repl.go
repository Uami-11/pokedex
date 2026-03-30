package main

import (
	"strings"
	"unicode"
)

func cleanInput(text string) []string {
	cleanOutput := []string{}
	trimText := strings.TrimSpace(text)
	var word strings.Builder
	for _, char := range trimText {
		if char == ' ' {
			if len(word.String()) > 0 {
				cleanOutput = append(cleanOutput, word.String())
				word.Reset()
			}
		} else {
			word.WriteRune(unicode.ToLower(char))
		}
	}

	cleanOutput = append(cleanOutput, word.String())
	return cleanOutput
}
