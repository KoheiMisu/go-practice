package main

import "fmt"

func main() {
	// 同じ配列データを共有している
	a := [3]int{1, 2, 3}
	s := a[:]
	fmt.Println(len(s))
	fmt.Println(cap(s))

	fmt.Println(s)

	// sは新たに確保されたメモリ領域を参照する
	// スライスは固定長を超えたときに2倍のcapをとる
	s = append(s, 4)
	fmt.Println(len(s))
	fmt.Println(cap(s))

	fmt.Println(s)
}
