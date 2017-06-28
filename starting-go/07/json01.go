package main

import (
	"encoding/json"
	"fmt"
	"log"
	"time"
)

type User struct {
	Id      int
	Name    string
	Email   string
	Created time.Time
}

func main() {
	u := new(User)
	u.Id = 1
	u.Name = "kohei"
	u.Email = "test@gmail.com"
	u.Created = time.Now()

	// json encode
	// json.Marshal return type is []byte
	// if debug we must use string(target)
	bs, err := json.Marshal(u)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(bs))

	// json.UnMarshal
	src := `
{
	"Created": "2017-06-28T23:43:20.271145492+09:00",
	"Email": "test@gmail.com",
	"Id": 1,
	"Name": "misu"
}
`
	u2 := new(User)

	// json decode
	err = json.Unmarshal([]byte(src), u2)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%+v\n", u2)
}
