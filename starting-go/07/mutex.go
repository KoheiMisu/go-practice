package main

import (
	"fmt"
	"sync"
	"time"
)

var st struct{ A, B, C int }

/**
sync.Mutex type can lock only one goroutine
this used to synchronize blocks of arbitrary processing

do not specify large range
*/
var mutex *sync.Mutex

func UpdatedAndPrint(n int) {
	// Lock
	mutex.Lock()

	// update each field of st with sleep

	// struct is mutable. so next asynchronous goroutine write another value.
	st.A = n
	time.Sleep(time.Microsecond)

	st.B = n
	time.Sleep(time.Microsecond)

	st.C = n
	time.Sleep(time.Microsecond)

	fmt.Println(st.A, st.B, st.C)

	mutex.Unlock()
}

func main() {
	// generate mutex
	mutex = new(sync.Mutex)

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
