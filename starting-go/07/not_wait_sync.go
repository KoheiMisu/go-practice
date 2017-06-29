package main

import (
	"fmt"
	"sync"
)

// mainの処理が実行終了になるとプログラム全体が終了する
// when main program finished, entire program finish.
// so this program dont output
func main() {
	go func() {
		for i := 0; i < 100; i++ {
			fmt.Println("1st Goroutine")
		}
	}()

	go func() {
		for i := 0; i < 100; i++ {
			fmt.Println("2nd Goroutine")
		}
	}()

	go func() {
		for i := 0; i < 100; i++ {
			fmt.Println("3rd Goroutine")
		}
	}()
}
