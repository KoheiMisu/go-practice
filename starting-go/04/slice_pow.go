package main

import "fmt"

// 参照型であるスライスで受ける
func pow(a []int) {
	/* スライスの各要素を2乗する */
	for i, v := range a {
		a[i] = v * v
	}
	return
}

func main() {
	/* 3要素の配列 */
	a := []int{1, 2, 3}
	// スライスは参照渡しである
	pow(a)
	fmt.Println(a) // => [1, 2, 3]
}
