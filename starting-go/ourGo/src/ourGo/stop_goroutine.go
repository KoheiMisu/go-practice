package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"sync"
	"time"
)

var wg sync.WaitGroup

var client http.Client

var re = regexp.MustCompile(`<title>.*</title>`)

// stop goroutine from outside
func main() {
	queue := make(chan string)

	for i := 0; i < 2; i++ { // generate 2 goroutine
		wg.Add(1)
		go fetchURL(queue)
	}

	queue <- "https://www.google.co.jp"
	queue <- "https://www.yahoo.co.jp/"
	queue <- "https://www.bing.com/"
	queue <- "https://www.microsoft.com/ja-jp/"

	close(queue) // notify goroutine of termination
	wg.Wait()    // wait until all goroutine finish
}

func fetchURL(queue <-chan string) {
	for {
		url, more := <-queue // if close, more will be false
		fmt.Println(more)
		if more {
			// get url
			fmt.Println("fetching url", url)
			resp, err := client.Get(url)

			if err != nil {
				// err
			}

			if resp.StatusCode == 200 { // OK
				bodyBytes, err2 := ioutil.ReadAll(resp.Body)
				if err2 != nil {

				}
				bodyString := string(bodyBytes)
				title := re.FindString(bodyString)
				fmt.Println(title)
			}
			time.Sleep(1 * time.Second)
		} else {
			fmt.Println("worker exit")
			wg.Done()
			return
		}
	}
}
