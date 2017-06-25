package main

import (
	"fmt"
	"math"
)

type Point struct {
	X, Y int
}

// Point型のメソッドRender
func (p *Point) Render() {
	fmt.Printf("<%d, %d>\n", p.X, p.Y)
}

// ↑ *Point型のpがレシーバーと呼ばれるもの
// funcと実装したいメソッドの間に置かないといけない

// 二点間の距離を求めるメソッド
func (p *Point) Distance(dp *Point) float64 {
	x, y := p.X-dp.X, p.Y-dp.Y
	return math.Sqrt(float64(x*x + y*y))
}

func main() {
	p := &Point{X: 5, Y: 12}
	// メソッドRenderの呼び出し
	p.Render()

	a := &Point{X: 6, Y: 9}
	a.Render()

	re := p.Distance(&Point{3, 4})
	fmt.Println(re)
}
