package main

import "fmt"

// 文字列化できることを示すインターフェース
type Stringify interface {
	ToString() string
}

// 構造体型Personの定義
type Person struct {
	Name string
	Age  int
}

func (p *Person) ToString() string {
	return fmt.Sprintf("%s(%d)", p.Name, p.Age)
}

// 構造体Carの定義
type Car struct {
	Number string
	Model  string
}

func (c *Car) ToString() string {
	return fmt.Sprintf("[%s] %s", c.Number, c.Model)
}

// Printlnの改良
func Println(s Stringify) {
	fmt.Println(s.ToString())
}

func main() {
	// 異なる型を共通するインターフェース型にまとめることができる
	vs := []Stringify{
		&Person{"Taro", 21},
		&Car{"000-000", "audi"},
	}

	for _, v := range vs {
		fmt.Println(v.ToString())
	}

	// Stringifyを実装している構造体にはPrintlnで内容を出力させられる
	Println(&Person{"hoge", 333})
}
