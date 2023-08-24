package main

import (
	"fmt"
	"strings"
)

func reverseWords(stringIn string) string {
	stringArray := strings.Fields(stringIn)
	for n, i := 0, len(stringArray)-1; n < i; n, i = n+1, i-1 {
		stringArray[n], stringArray[i] = stringArray[i], stringArray[n]
	}
	return strings.Join(stringArray, " ")
}

func main() {
	fmt.Printf("Output: %q\n", reverseWords("hello world"))
	fmt.Printf("Output: %q\n", reverseWords("  I love programming "))
	fmt.Printf("Output: %q\n", reverseWords("  the sky is blue  "))
}
