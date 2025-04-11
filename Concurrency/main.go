package main

import (
	"fmt"
	"time"
)

func greet(phrase string, done chan bool) {
	fmt.Println("Hello!", phrase)
	done <- true
}

func slowGreet(phrase string, done chan bool) {
	time.Sleep(3 * time.Second) // simulate a slow, long-taking task
	fmt.Println("Hello!", phrase)
	done <- true
	//close(done) only makes sense if the channel is used for single-use signals and we know that the channel will not be used again and it is the long-running task
}

func main() {
	dones := make([]chan bool, 4)

	for i := 0; i < 4; i++ {
		dones[i] = make(chan bool)
	}

	go greet("Nice to meet you!", dones[0])
	go greet("How are you?", dones[1])
	go slowGreet("How ... are ... you ...?", dones[2])
	go greet("I hope you're liking the course!", dones[3])

	for _, done := range dones {
		<-done
	}

	fmt.Println("Done")
}
