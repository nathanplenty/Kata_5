package main

import (
	"fmt"
	"strings"
)

func reverseWords(stringIn string) string {
	// Get each Input word without spacers as array for strings
	stringArray := strings.Fields(stringIn)
	// Initialize empty output 'stringOut'
	stringOut := ""
	// Append each word from 'stringArray' backwards on 'stringOut'
	for i := len(stringArray); i > 0; i-- {
		stringOut += stringArray[i-1]
		if i <= len(stringArray) {
			stringOut += " "
		}
	}
	// Remove last spacer from 'stringOut'
	stringOut = strings.TrimRight(stringOut, " ")
	return stringOut
}

func main() {
	fmt.Printf("Output: %q\n", reverseWords("hello world"))
	fmt.Printf("Output: %q\n", reverseWords("  I love programming "))
	fmt.Printf("Output: %q\n", reverseWords("  the sky is blue  "))
}
