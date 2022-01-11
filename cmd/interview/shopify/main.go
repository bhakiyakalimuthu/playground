package main

import (
	"flag"
	"fmt"
	"strconv"
	"strings"
)

const (
	pass = "PASS"
	fail = "INVESTIGATE"

	number = iota
	letter
	unknown
)

func main() {
	//CAJI202002196 # ID code
	//202002196 # Numeric part
	//20200219 # Ignoring verification digit
	//2 + 2 + 0 + 1 = 5 # Sum of odd position digits, i.e. O
	//0 + 0 + 2 + 9 = 11 # Sum of even position digits, i.e. E
	//5 - 11 = -6 # V = O - E
	//6 # Verification digit, ignoring sign

	//validateCode()
	display([]int{1, 6, 45, 35, 20, 39, 19})
}
func display(prices []int) {
	if len(prices) == 0 {
		fmt.Println("empty price list")
	}
	s0 := make([]string, 0)
	s0 = append(s0, "00:")
	s1 := make([]string, 0)
	s1 = append(s1, "10:")
	s2 := make([]string, 0)
	s2 = append(s2, "20:")
	s3 := make([]string, 0)
	s3 = append(s3, "30:")
	s4 := make([]string, 0)
	s4 = append(s4, "40:")

	for _, val := range prices {
		switch {
		case val <= 9 && val >= 0:
			s0 = append(s0, "*")
		case val <= 19 && val >= 10:
			s1 = append(s1, "*")
		case val <= 29 && val >= 20:
			s2 = append(s2, "*")
		case val <= 39 && val >= 30:
			s3 = append(s3, "*")
		case val <= 49 && val >= 40:
			s4 = append(s4, "*")
		default:
			fmt.Println("unknown")
		}
	}
	//fmt.Println(s0, s1)
	fmt.Printf("%v\n", strings.Trim(fmt.Sprint(s0), "[]"))
	fmt.Printf("%v\n", strings.Trim(fmt.Sprint(s1), "[]"))
	fmt.Printf("%v\n", strings.Trim(fmt.Sprint(s2), "[]"))
	fmt.Printf("%v\n", strings.Trim(fmt.Sprint(s3), "[]"))
	fmt.Printf("%v\n", strings.Trim(fmt.Sprint(s4), "[]"))
	//fmt.Printf("%v\n%v\n%v\n%v\n%v\n", strings.Trim(fmt.Sprint(s0, "\n", s1, s3), "[]"), s1, s2, s3, s4)
}
func validateCode() string {
	fName := flag.String("f", "first name", "Enter your first name")
	lName := flag.String("l", "last name", "Enter your last name")
	code := flag.String("c", "CODE", "Enter your code")
	flag.Parse()
	//if len(*code) != 13 {
	//	fmt.Println("INVESTIGATE")
	//	return fail
	//}
	fmt.Printf("First name %s\n", *fName)
	fmt.Printf("last name %s\n", *lName)
	fmt.Printf("code %d\n", *code)
	//str := "abcdefghijklmnopqrstuvwxyz"
	//str := "ABCDEFGHIJKLMNOPQRSTUVWXYZ123ajflkdjalfdjlf"
	//str := "0123456789"

	for i, char := range (*code)[:4] {
		if checkType(char) == number {
			return fail
		}
		fmt.Println(i, char)
	}
	for i, char := range (*code)[4:] {
		if checkType(char) != number {
			return fail
		}
		num, _ := strconv.Atoi(string(char))
		fmt.Println(num)
		fmt.Println(i, char)
	}
	return ""
}
func checkType(char rune) int {
	switch {
	case char >= '0' && char <= '9':
		fmt.Printf("char : %s is number\n", string(char))
		return number
	case char >= 'a' && char <= 'z' || char >= 'A' && char <= 'Z':
		fmt.Printf("char : %s is a-z\n", string(char))
		return letter
	default:
		return unknown
	}
}
