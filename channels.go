package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	sites := []string{
		"http://google.com", "http://facebook.com", "http://stackoverflow.com", "http://golang.org", "http://amazon.com",
	}

	//channel
	c := make(chan string)

	for _, site := range sites {
		// go keyword only in front of function calls
		go checkLink(site, c)
	}
	//Reading messages from channel is blocking call similar to future.get() in java
	// fmt.Println(<-c)
	// fmt.Println(<-c)
	// fmt.Println(<-c)
	// fmt.Println(<-c)
	// fmt.Println(<-c)
	//Since only 5 message come into the channel 6th call will never end, so program hangs
	// fmt.Println(<-c)
	// for i := 0; i < len(sites); i++ {
	// 	fmt.Println(<-c)
	// }
	// for {
	// 	go checkLink(<-c, c)
	// }

	//Iterate over messages in the channel
	for s := range c {
		//function literal in go -- similar to lambda function or anonymous fucntion in java
		//Pass variable to go routine with method arg, so that pass by value copies the variable
		//do not reference the main go routine varaible directly in child co routine
		go func(site string) {
			time.Sleep(5 * time.Second)
			checkLink(site, c)
		}(s)
	}
}

func checkLink(site string, c chan string) {
	res, err := http.Get(site)
	if err != nil {
		fmt.Println(site, " is down: ", err)

		c <- site
	} else {
		fmt.Println(site, res.Status)

		c <- site
	}
}
