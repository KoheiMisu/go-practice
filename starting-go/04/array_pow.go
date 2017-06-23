package main

import (
	"fmt"
)

// 配列は値渡し
func pow(a [3]int) {
	/* 配列の各要素を2乗する */
	for i, v := range a {
		a[i] = v * v
	}
	return
}

func main() {
	/* 3要素の配列 */
	a := [3]int{1, 2, 3}
	pow(a)
	fmt.Println(a) // => [1, 2, 3]
}
