package main

import (
	"fmt"
	"log"
	"os"
	"time"
)

func main() {
	// os.Exitによって破棄されるdefer
	defer func() {
		fmt.Println("defer")
	}()

	fmt.Println(time.January)
	fmt.Println(time.January.String())

	_, err := os.Open("foo!!!")
	if err != nil {
		log.Fatal(err)
	}

	// プログラムの終了
	// 上記のdeferは実行されない
	//os.Exit(0) //終了ステータス0で終了
}
