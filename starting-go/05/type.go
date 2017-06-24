package main

import "fmt"

// int型をとり、int型を返すコールバック型
type Callback func(i int) int

func Sum(ints []int, callback Callback) int {
	var sum int
	for _, i := range ints {
		sum += i
	}

	return callback(sum)
}

func main() {
	// type [定義する型] [既存の型]

	/**
	int型にMyIntのようなaliasを置くこともできる
	*/
	type MyInt int
	var nl MyInt = 5

	n2 := MyInt(7)
	fmt.Println(nl)
	fmt.Println(n2)

	// typeによるさまざまなエイリアス
	type (
		IntPair     [2]int
		strings     []string
		AreaMap     map[string][2]float64 // 地域名と緯度・経度を保持
		IntsChannel chan []int
	)

	n := Sum(
		[]int{1, 2, 3, 4, 5},
		func(i int) int {
			return i * 2
		}, // 最後の引数はカンマでトレイリングしていないといけない
	)

	fmt.Println(n) // => 30

}
