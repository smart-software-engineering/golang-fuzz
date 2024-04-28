package main

import (
	"errors"
	"fmt"
	"unicode/utf8"
)

func Reverse(s string) (string, error) {
	if !utf8.ValidString(s) {
		return s, errors.New("input is not valid UTF-8")
	}
	b := []rune(s)
	for i, j := 0, len(b)-1; i < len(b)/2; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}
	return string(b), nil
}

func main() {
	input := "The quick brown fox jumped over the lazy dog"
	rev, revError := Reverse(input)
	doubleRev, doubleRevError := Reverse(rev)
	fmt.Printf("original: %q\n", input)
	fmt.Printf("reversed: %q, err: %v\n", rev, revError)
	fmt.Printf("reversed again: %q, err: %v\n", doubleRev, doubleRevError)
}
