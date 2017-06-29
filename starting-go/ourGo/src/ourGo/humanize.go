package main

import (
	"fmt"
	"github.com/dustin/go-humanize"
)

func main() {
	fmt.Printf("That file is %s.", humanize.Bytes(82854982)) // => That file is 83 MB.
}
