package main

import "fmt"

type Myint int

type Strings []string

func (m Myint) Plus(i int) int {
	return int(m) + i
}

func (s Strings) Join(d string) string {
	sum := ""
	for _, v := range s {
		if sum != "" {
			sum += d
		}
		sum += v
	}
	return sum
}

func main() {
	Myint(4).Plus(2) // == 6

	Strings{"S", "O", "S"}.Join(",") // == SOS
}
