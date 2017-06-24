package main

import (
	"fmt"
)

func main() {
	m := map[int]string{1: "A", 2: "B", 3: "C"}

	s2, ok := m[100]
	fmt.Println(ok) // => false
	fmt.Println(s2)

	/* 要素が存在したら行う処理 */
	if _, ok := m[1]; ok {
		fmt.Println(m[1])
	}

	/* for */
	for key, val := range m {
		fmt.Printf("%d => %s\n", key, val)
	}

	// len()は使用できるが、capはない
	fmt.Println(len(m))

	// 要素を削除する時はdeleteを使う
	delete(m, 2)
	fmt.Println(m)
}
