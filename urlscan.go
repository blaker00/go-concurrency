package main

import (
	"fmt"
	"net/http"
)

func grabUrl(ch1 chan string, url string) <-chan string {
	_, err := http.Get(url)
	if err != nil {
		fmt.Printf("There was an error getting url: %s\n", url)
	}
	ch1 <- fmt.Sprintf("URL: %s, was successful", url)

	return ch1
}

func main() {
	ch1 := make(chan string, 2)

	urls := []string{
		"https://www.yahoo.com",
		"https://www.google.com",
		"https://www.slashdot.com",
		"https://www.youtube.com",
		"https://www.facebook.com",
	}

	go func() {
		for _, v := range urls {
			grabUrl(ch1, v)
		}
		close(ch1)
	}()

	for v := range ch1 {
		fmt.Println(v)
	}

}
