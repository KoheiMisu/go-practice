package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Id   int    "ID"
	Name string "名前"
}

func main() {
	u := User{1, "author"}

	/* 変数tはreflect.Type型 */
	t := reflect.TypeOf(u)

	/* 構造体の全フィールドを処理するループ */
	for i := 0; i < t.NumField(); i++ {
		/* i番目のフィールドを取得 */
		f := t.Field(i)
		fmt.Println(f.Name, f.Tag)
	}
}
