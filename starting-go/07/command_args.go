package main

import (
	"fmt"
	"os"
)

// ビルドしてから行う
// getargs [arg] [arg] [arg]
func main() {
	fmt.Println(os.Args[0]) // => getargs
	fmt.Println(os.Args[1])
	fmt.Println(os.Args[2])
	fmt.Println(os.Args[3])
}
