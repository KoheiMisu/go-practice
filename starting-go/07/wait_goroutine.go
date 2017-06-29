package main

import (
	"fmt"
	"sync"
)

// sync.WaitGroup型は、 任意のゴルーチンによる処理の完了を待ち受けるための仕組みを提供
func main() {
	// generate sync.WaitGroup
	wg := new(sync.WaitGroup)

	// declare count for goroutine
	wg.Add(3)

	go func() {
		for i := 0; i < 100; i++ {
			fmt.Println("1st Goroutine")
		}
		// finished
		wg.Done()
	}()

	go func() {
		for i := 0; i < 100; i++ {
			fmt.Println("2nd Goroutine")
		}
		// finished
		wg.Done()
	}()

	go func() {
		for i := 0; i < 100; i++ {
			fmt.Println("3rd Goroutine")
		}
		// finished
		wg.Done()
	}()

	// wait finished goroutine
	wg.Wait()
}
