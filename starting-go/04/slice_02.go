package main

import (
	"fmt"
)

// 可変長引数をとる関数. ただし、複数は取れない。引数の末尾に１つのみ
func sum(s ...int) int {
	n := 0

	for _, v := range s {
		n += v
	}

	return n
}

func main() {
	s := []string{"Apple", "Banana", "Cherry"}

	// foreach
	for i, v := range s {
		fmt.Printf("[%d] => %s\n", i, v)
	}

	fmt.Println(sum(1, 2, 3)) // => 6
	fmt.Println(sum())        // => 0

	t := []int{1, 2, 3}
	// スライスを可変長引数として扱う
	fmt.Println(sum(t...)) // => 6
}
