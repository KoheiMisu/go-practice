package main

import "fmt"

// error型は、文字列を返すErrorメソッドのみが定義されている
//type error interface {
//	Error() string
//}

// 独自定義のエラーを表す型
type MyError struct {
	Message string
	Errcode int
}

func (e *MyError) Error() string {
	return e.Message
}

func RaiseError() error {
	return &MyError{
		Message: "エラーが発生しました",
		Errcode: 1234,
	}
}

type T struct {
	Id   int
	Name string
}

// fmt.Stringer
//type Stringer interface {
//	String() string
//}

func (t *T) String() string {
	return fmt.Sprintf("<%d %s>", t.Id, t.Name)
}

func main() {
	err := RaiseError()
	fmt.Println(err.Error())

	// 型アサーションによって本来の型を取り出す
	e, ok := err.(*MyError)
	if ok {
		fmt.Println(e.Errcode)
	}

	t := &T{10, "hoge"}
	fmt.Println(t) // => &{10, "hoge"} // String()インターフェース実装前の場合

	t2 := &T{11, "hogehoge"}
	fmt.Println(t2)
}
