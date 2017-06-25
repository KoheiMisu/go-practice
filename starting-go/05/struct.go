package main

import (
	"fmt"
)

type Point struct {
	X int // 初期値は型に合わせた初期値。ここではint型なので0
	Y int
}

// constructor: 一般的にはNew + 構造体名
func NewPoint(x, y int) Point {
	pt := Point{x, y}
	return pt
}

//type Point struct {
//	X, Y int
//}

type Feed struct {
	Name   string
	Amount uint
}

type Animal struct {
	Name string
	Feed /* 省略しておくとanimal.Amountでアクセスできる */
}

func swap(p *Point) {
	x, y := p.Y, p.X
	p.X = x
	p.Y = y
}

func main() {
	var pt Point
	pt.X = 1
	pt.Y = 2

	fmt.Println(pt)

	pt2 := Point{3, 4}
	fmt.Println(pt2)

	pt3 := Point{X: 5, Y: 6}
	fmt.Println(pt3)

	pt4 := NewPoint(7, 8)
	fmt.Println(pt4)

	a := Animal{
		Name: "Monkey",
		Feed: Feed{
			Name:   "Banana",
			Amount: 10,
		},
	}

	fmt.Println(a)
	fmt.Println(a.Feed.Name)
	fmt.Println(a.Amount)

	//a2 := Animal{
	//	Name: "Bird",
	//	Feed: 1, // Feedの型省略をした場合、自然にFeed型を参照するように縛れる
	//}
	// ↑ Error, cannot use 1 (type int) as type Feed in field value

	//fmt.Println(a2)

	// 構造体は値型なので、フィールドを操作する場合はポインタを渡すようにする
	// 通常、構造体はポインタを直接生成する
	p5 := &Point{10, 11}
	swap(p5)
	fmt.Println(p5) // => 11, 10
}
