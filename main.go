package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://amazon.com",
		"http://netflix.com",
		"http://seeingblue.us",
	}

	c := make(chan string)

	for _, link := range links {
		go checkLink(link, c)
	}

	//for i := 0; i < len(links); i++ {
	//for {
	for l := range c {
		go func(link string) {
			time.Sleep(5 * time.Second)
			checkLink(link, c)
		}(l)
	}
}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {

		fmt.Println(link + " is Down")
		c <- link
		return
	}
	fmt.Println(link + " is Up")
	c <- link
}
