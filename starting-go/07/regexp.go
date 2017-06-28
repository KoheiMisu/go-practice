package main

import (
	"fmt"
	"regexp"
)

func main() {
	// not need escape when we use raw literal. ex) \\d -> \d
	re := regexp.MustCompile(`\d+`)
	fmt.Println(re.MatchString("44")) // => true

	// extract matched parts
	extracted := re.FindString("xxx888yyy999")
	fmt.Println(extracted)

	// select matched part and extract
	extracted2 := re.FindAllString("xxx888yyy999", 1)
	fmt.Println(extracted2) // => [888]

	extracted3 := re.FindAllString("xxx888yyy999", -1)
	fmt.Println(extracted3) // => [888 999]

	// split strings by regexp
	rege01 := regexp.MustCompile(`\s+`)
	re01 := rege01.Split("a b c d", -1)
	fmt.Println(re01) // => [a b c d]

	// replace by regexp
	rege02 := regexp.MustCompile(`\d+`)
	re02 := rege02.ReplaceAllString("apple is 100 yen", "2000")
	fmt.Println(re02)

	// extract grouping
	s3 := `
00-111-1111
080-0000-0000
`
	rege03 := regexp.MustCompile(`(\d+)-(\d+)-(\d+)`)
	re03 := rege03.FindAllStringSubmatch(s3, -1)
	for _, v := range re03 {
		fmt.Println(v)
	}
	// => [00-111-1111 00 111 1111]
	// => [080-0000-0000 080 0000 0000]

	// group and replace by regexp
	fmt.Println(rege03.ReplaceAllString("Tel: 000-111-222", "$3-$2-$1")) // Tel: 222-111-000

}
