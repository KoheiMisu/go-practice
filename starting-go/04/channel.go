package main

import (
	"fmt"
)

func main() {
	// チャネル chan型

	/* 変数chは、int型のチャネル */
	var ch chan int

	/* int型の受信専用チャンネル */
	var ch1 <-chan int

	/* int型の送信専用チャンネル */
	var ch2 chan<- int

	// バッファサイズ0のチャネル(バッファサイズが指定されない場合は0)
	ch3 := make(chan int)

	// バッファサイズ8で初期化
	ch4 := make(chan int, 8)

	/**
	 * チャネルは、キュー(待ち行列)の性質を備えるデータ構造
	 * FIFO(先入れ先出し)
	 */
}
