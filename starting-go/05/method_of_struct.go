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
// レシーバーは原則ポインタ型で設定
// 値型で受けるとコピーが発生し、構造体のフィールドをいじるような処理は意図した動きにならない

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

	// 構造体型をマップのキーにするにせよ、値にするにせよ
	// リテラル内で構造体の型名を省略できる. => Point{1, 2}: "author", のようにしなくてよい
	m1 := map[Point]string{
		{1, 2}: "author",
	}

	fmt.Println(m1)
}
