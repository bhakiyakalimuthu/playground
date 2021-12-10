package main

import (
	"fmt"
)

func main() {
	// create stack
	// if open braces, push it to stack
	// if closed braces, pop it out from the stack,if same type
	// if not same type return false
	// once the loop is closed if the stack is not empty return false
	fmt.Println(isBalanced("{{(())}}"))
}

func isBalanced(in string) bool {
	if len(in) == 0 {
		return false
	}
	m := map[rune]rune{
		'}': '{',
		')': '(',
		']': '[',
	}
	s := make([]rune, 0, len(in))
	for _, v := range in {
		switch v {
		case '{':
			s = append(s, v)
		case '(':
			s = append(s, v)
		case '[':
			s = append(s, v)
		default:
			val, ok := m[v]
			if len(s) == 0 || !ok || val != s[len(s)-1] {
				return false
			}
			s = s[:len(s)-1]
		}
	}
	return len(s) == 0

}
