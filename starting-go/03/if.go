package main

import "fmt"

func doSomething() bool {
	return false
}

/**
簡易文付きif
*/
func main() {
	if x, y := 1, 2; x < y {
		fmt.Printf("x(%d) is less than y(%d)\n", x, y)
	}

	z := 5
	if z := 3; true {
		fmt.Printf("z(%d)\n", z)
	}

	fmt.Printf("z(%d)\n", z)

	a, b := 3, 5

	if n := a * b; n%2 == 0 {
		fmt.Printf("n(%d) is even \n", n)
	} else {
		fmt.Printf("n(%d) is odd \n", n)
	}

	//if _, err := doSomething(); err != nil {
	//	fmt.Println("something is wrong")
	//}
}
