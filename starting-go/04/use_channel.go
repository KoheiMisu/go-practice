package main

import "fmt"

func main() {
	ch := make(chan int, 10)

	ch <- 7
	ch <- 5
	ch <- 15
	i := <-ch
	i2 := <-ch

	fmt.Println(i)  // => 7
	fmt.Println(i2) // => 5

	// channelはclosedという状態を持つ
	ch2 := make(chan int, 1)
	close(ch2)
	// closeしたチャンネルに送信はできない
	//ch2 <- 1

	// 受信の処理はチャネルが内包する型の初期値を受信する
	t := <-ch2
	fmt.Println(t) // => 0

	// チャネルがクローズされているか調べる
	ch3 := make(chan int)
	close(ch3)

	i, ok := <-ch // i ==0, ok == false

	// チャネルに残存するデータがある場合はtrueを返す

}
