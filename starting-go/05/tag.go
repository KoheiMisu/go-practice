package main

import (
	"encoding/json"
	"fmt"
	"reflect"
)

type User struct {
	Id   int    "ID"
	Name string "名前"
}

type Author struct {
	Id   int    `json:"author_id"` // jsonパッケージが出力時のキーを拾うように設定
	Name string `json:"author_name"`
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

	a := Author{999, "999"}
	bs, _ := json.Marshal(a)
	fmt.Println(string(bs))
	fmt.Printf("%s\n", bs)
}
