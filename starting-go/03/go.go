package main

import (
	"fmt"
)

func sub() {
	for {
		fmt.Println("sub loop")
	}
}

func main() {
	go sub() // ゴルーチン開始

	for {
		fmt.Println("main loop")
	}
}
