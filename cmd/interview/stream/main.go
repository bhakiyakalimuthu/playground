package main

import "fmt"

// SquareInPlace replaces the value of v with its square value
// Note: this is a void function
// Note: square of two is four and not 1.41421356237...
// TODO: replace T with the appropriate type
func SquareInPlace(v *float64) {
	// TODO: implement the square in place for v
	*v = *v * *v
}

func main() {
	x := 1.5
	// TODO: replace the ? placeholder
	SquareInPlace(&x)
	fmt.Print(x == 2.25)
}
