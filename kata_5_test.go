package main

import "testing"

func TestYes(t *testing.T) {
	b := false
	sw := "world hello"
	sg := reverseWords(sw)
	if sw == sg {
		b = true
	}
	if b == true {
		t.Errorf("b = %q; want 'world hello'", sg)
	}
}

func TestNo(t *testing.T) {
	b := true
	sw := "hello world"
	sg := reverseWords(sw)
	if sw != sg {
		b = false
	}
	if b == true {
		t.Errorf("b = %v; want 'hello world'", sg)
	}
}
