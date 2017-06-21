package main

import (
	"fmt"
)

/* nはパッケージ変数(global) */
var n = 100

func main() {
	n = n + 1
	fmt.Printf("n = %d\n", n)
}
