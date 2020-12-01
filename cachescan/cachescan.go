package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	// Do not communicate by sharing memory, share memory by communicating
	fmt.Println("vim-go")

	start := time.Now()
	//ch := make(chan string)
	HTTPGet("https://news.ycombinator.com")
	HTTPGet("https://www.france24.com")
	HTTPGet("https://www.jreast.co.jp/")
	HTTPGet("https://www.flashback.org/")
	fmt.Printf("Total elapsed\t%.3fs\n", time.Since(start).Seconds())
}

func HTTPGet(url string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
	}
	// The deferred function call is scheduled just before HTTPGet returns
	defer resp.Body.Close()
	fmt.Printf("%s\t%.3fs\n", url, time.Since(start).Seconds())

}
