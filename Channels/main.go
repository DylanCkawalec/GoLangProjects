package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	links := []string{
		"http://google.com",
		"http://stackoverflow.com",
		"http://game.defikingdoms.com/#/"}

	c := make(chan string)

	for _, link := range links {
		go errorCheckLink(link, c)
	}
	//this is passing a channel to a value of the sametype, and being passed to l, and c must be called to produce a string.
	for l := range c {
		//function literall to avoid backup, which is an Anonymous funciton
		go func(link string) {
			time.Sleep(5 * time.Second)
			errorCheckLink(link, c)
		}(l)
		//this passes the link to a child routine as a copy in memory, pass by value

	}

}

func errorCheckLink(link string, c chan string) {
	_, err := http.Get(link) //struct and err returns. we only care about error however

	if err != nil {
		fmt.Println(link, "might be down!")
		c <- link
		return
	}

	fmt.Println(link, "is up!")
	c <- link
}
