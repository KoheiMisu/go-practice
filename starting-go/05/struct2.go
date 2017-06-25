package main

import "fmt"

type Base struct {
	Id    int
	Owner string
}

type A struct {
	Base /* 共通のフィールド */
	Name string
	Area string
}

type B struct {
	Base
	Title  string
	Bodies []string
}

type Person struct {
	Id   int
	Name string
	Area string
}

type User struct {
	Id   int
	Name string
}

func NewUser(id int, name string) *User {
	u := new(User)
	u.Id = id
	u.Name = name
	return u
}

// 再帰的な定義の禁止
// 自身の型を含むようなstructは作成できない
//type T struct {
//	T
//}

func main() {
	a := A{
		Base: Base{9, "me"},
		Name: "hoge",
		Area: "tokyo",
	}

	fmt.Println(a)

	// 指定した型のポインタ型を生成するための組み込み関数 -> new
	p := new(Person)
	fmt.Println(p) // => &{0 "" ""}
}
