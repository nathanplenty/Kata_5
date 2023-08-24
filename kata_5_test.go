package main

import "testing"

func TestYes(t *testing.T) {
	boolean := false
	stringTarget := "world hello"
	stringIn := reverseWords(stringTarget)
	if stringTarget == stringIn {
		boolean = true
	}
	if boolean == true {
		t.Errorf("boolean = %q; want 'world hello'", stringIn)
	}
}

func TestNo(t *testing.T) {
	boolean := true
	stringTarget := "hello world"
	stringIn := reverseWords(stringTarget)
	if stringTarget != stringIn {
		boolean = false
	}
	if boolean == true {
		t.Errorf("boolean = %v; want 'hello world'", stringIn)
	}
}
