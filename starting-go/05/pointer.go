package main

import "fmt"

func inc(a *int) {
	// aをデリファレンスして+1の処理を行う
	*a++
}

func printString(s string) {
	fmt.Println(s)
}

func main() {
	var p *int

	fmt.Println(p == nil) // => true

	var i int
	p2 := &i
	i = 5
	fmt.Println(*p2)
	// デリファレンスでポインタ型が指し示すデータを参照
	*p2 = 10
	fmt.Println(i)

	a := 1
	inc(&a)
	inc(&a)
	inc(&a)
	fmt.Println(a)

	// ポインタ型の変数がnilである場合にデリファレンスを実行するとランタイムパニックが発生
	var p *int
	fmt.Println(*p)

	// 文字列のポインタ扱いについて
	s := "ABC"
	&s    // 文字列の型のポインタ
	s[0]  // 文字列のインデックス参照(byte型)
	&s[0] // コンパイルエラー

	/**
	 Goでは一度生成された文字列に対して何らかの変更を加えることは
	 基本的にできないようになっている
	 → immutableとして扱う設計になっている

	 string型は内部的に
	   - 文字列の実態へのポインタ
	   - 文字列のバイト長
	を持っている
	*/

	st := "pointerrrrrr!!!"
	printString(st) // 文字列の実態のコピーは行われない。メモリ位置を参照する
}
