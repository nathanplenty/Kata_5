package main

import (
	"fmt"
	"strings"
)

func reverseWords(s string) string {
	// Get each Input word without spacers as array for strings
	stringArray := strings.Fields(s)
	// Clear Input 's'
	s = ""
	// Append each word from 'stringArray' backwards on 's'
	for i := len(stringArray); i > 0; i-- {
		s += stringArray[i-1]
		if i <= len(stringArray) {
			s += " "
		}
	}
	// Remove last spacer from 's'
	s = strings.TrimRight(s, " ")
	return s
}

func main() {
	fmt.Printf("Output: %q\n", reverseWords("hello world"))
	fmt.Printf("Output: %q\n", reverseWords("  I love programming "))
	fmt.Printf("Output: %q\n", reverseWords("  the sky is blue  "))
}
