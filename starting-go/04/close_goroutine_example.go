package main

import (
	"fmt"
	"time" // ウエイトのために利用
)

func receive(name string, ch <-chan int) {
	for {
		i, ok := <-ch
		if ok == false {
			// 受信できなくなったら終了
			break
		}

		fmt.Println(name, i)
	}
	fmt.Println(name + "is done.")
}

func main() {
	ch := make(chan int, 20)

	go receive("1st gorouitne", ch)
	go receive("2nd gorouitne", ch)
	go receive("3rd gorouitne", ch)

	i := 0
	for i < 100 {
		ch <- i
		i++
	}

	close(ch)

	time.Sleep(3 * time.Second)
}
