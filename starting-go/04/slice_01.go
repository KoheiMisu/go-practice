package main

import (
	"fmt"
)

func main() {
	s := make([]int, 8)
	fmt.Printf("%d", len(s))

	/**
	 * 第三引数を宣言すると容量であるcapacityを設定できる
	 * スライスのcapを超えた場合はメモリ領域の確保と新しい領域へのコピーが
	 * 行われるのでコストが高い
	 */
	sl := make([]int, 5, 10)
	fmt.Printf("%d", len(sl))
	fmt.Printf("%d", cap(sl))

	// 簡易スライス
	slice := [5]int{1, 2, 3, 4, 5}
	sliced := slice[0:2]
	fmt.Println(sliced)
	fmt.Println(slice[:])  // index 0 ~ len(slice)-1 までのスライス
	fmt.Println(slice[:2]) // index 0 ~ 2-1 までのスライス

	// append
	// 変数の代入を伴わないものはエラーになる ex.) append(sl, 10, 11, 12)
	// 配列の追加なので同じ変数で受けなければならない?
	s2 := []int{1, 2, 3}
	s2 = append(s2, 1)
	fmt.Println(s2)

	// copy
	// 返り値はコピーが実行された要素数
	t1 := []int{1, 2, 3, 4, 5}
	t2 := []int{10, 11}
	t3 := copy(t1, t2)
	fmt.Println(t3) // => 2, t1 == [10, 11, 3, 4, 5]

	// 完全スライス式
	// a[low : high : max] ( 0 <= low <= high <= max <= cap(a) )
	a := [10]int{}
	a1 := a[2:4]   // => index2 ~ 9まで値が格納可能。容量8
	a2 := a[2:4:6] // => index2 ~ 5まで値が格納可能。容量4(0~5まで容量とるが、indexが2からなので)

}
