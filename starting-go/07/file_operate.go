package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	f, err := os.Open("../inputFiles/hoge.txt")
	if err != nil {
		log.Fatal(err)
	}

	// ファイルの情報取得
	fi, err := f.Stat()
	fmt.Println(fi.Name())
	fmt.Println(fi.Size())
	fmt.Println(fi.Mode())
	fmt.Println(fi.ModTime())
	fmt.Println(fi.IsDir())

	// ファイル作成(指定したファイルが既に存在する場合は、ファイルの内容が削除されサイズ0になる)
	cf, err := os.Create("bar.txt")

	// 新規作成したファイルのステータス
	fi2, _ := cf.Stat()
	fmt.Println(fi2.Name())
	fmt.Println(fi2.Size())
	fmt.Println(fi2.IsDir())

	// ファイルに[]byte型の内容を書き込み
	cf.Write([]byte("Hello World\n"))
	//　オフセットを指定して書き込み
	cf.WriteAt([]byte("Golang"), 5)
	// ファイルの末尾にオフセットを移動
	cf.Seek(0, os.SEEK_END)
	// 文字列をファイルに書き込み
	cf.WriteString("Yeah!!!")

	// 関数終了時に確実にクローズする
	defer f.Close()
}
