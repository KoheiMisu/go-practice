package main

import (
	"bufio"
	"fmt"
	"strings"
)

func main() {
	s := `test
XXXXXXXXXX
YYYYYYYYYY
ZZZZZZZZZZ
`

	// generate Reader from string
	r := strings.NewReader(s)

	scanner := bufio.NewScanner(r)

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	// switch scanner scan method
	s2 := `golang PHP JAVA TuLLYs`

	// scanner will target rows by default
	r2 := strings.NewReader(s2)
	scanner2 := bufio.NewScanner(r2)

	/**
	to change the behavior of the scanner,
	pass another scan function defined in bufio
	*/

	// change scan function to bufio.ScanWords
	scanner2.Split(bufio.ScanWords)

	for scanner2.Scan() {
		fmt.Println(scanner2.Text())
	}
}
