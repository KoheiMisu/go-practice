package main

import "fmt"

func main() {
	s := "A"

	// switchはデフォでcaseに当てはまった処理をした後抜ける
	// 下にも流す場合はfallthroughを使う
	switch s {
	case "A":
		s += "B"
		fallthrough
	case "B":
		s += "C"
	}

	fmt.Printf("%s", s)
}
