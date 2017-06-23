package main

import "fmt"

var (
	a [3]int    // => [0, 0, 0] 要素の型の初期値によって内容が埋め込まれる
	b [3]string // => ['', '', ''] 要素の型の初期値によって内容が埋め込まれる
	// 参照型とは何らかの値を指す参照を保持しているデータ型
	c []int // nilが初期値になる → 値への参照を保持していないから
)

func main() {
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c == nil)
}
