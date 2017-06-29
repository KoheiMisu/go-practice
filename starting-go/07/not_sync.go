package main

import (
	"fmt"
	"time"
)

var st struct{ A, B, C int }

func UpdatedAndPrint(n int) {
	// update each field of st with sleep

	// struct is mutable. so next asynchronous goroutine write another value.
	// this called race condition
	st.A = n
	time.Sleep(time.Microsecond)

	st.B = n
	time.Sleep(time.Microsecond)

	st.C = n
	time.Sleep(time.Microsecond)

	fmt.Println(st.A, st.B, st.C)
}

func main() {
	// boot multiple goroutine
	for i := 0; i < 5; i++ {
		go func() {
			for i := 0; i < 1000; i++ {
				UpdatedAndPrint(i)
			}
		}()
	}
	for {
	}
}
