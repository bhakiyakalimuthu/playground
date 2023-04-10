package main

func namedReturnCallee(x int) (result int) {
	result = 10
	if x > 0 {
		return // this has the effect of " return 10"
	}
	go func() {
		_ = result // read result
	}()
	return 20 // this is equivalent to result =20
}

func main() {
	_ = namedReturnCallee(0)
}
