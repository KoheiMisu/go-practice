package main

import (
	"fmt"
)

func main() {
	// int型のキーとstring型の値を保持するマップ
	// var m map[int]string

	m := make(map[int]string)

	m[1] = "US"
	m[81] = "Japan"

	fmt.Println(m)

	// キーがint型、要素が[]int型であるマップ
	m2 := map[int][]int{
		1: {1},       // []int{1}
		2: {1, 2},    // []int{1, 2}
		3: {1, 2, 3}, // []int{1, 2, 3}  ← 冗長(Redundant)と言われたので止めた
	}

	fmt.Println(m2)

}
