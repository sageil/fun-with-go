package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.NewSource(time.Now().UnixNano())

	ch1 := make(chan string)
	ch2 := make(chan string)
	go func() {
		rnd := rand.Intn(1000) + 100
		for i := 0; i < rnd; i++ {
			ch1 <- fmt.Sprintf("Hello #%d", i)
		}
	}()
	go func() {
		rnd := rand.Intn(1000) + 100
		for i := 0; i < rnd; i++ {
			ch2 <- fmt.Sprintf("Salut #%d", i)
		}
	}()

	go func() {
		defer close(ch1)
		defer close(ch2)
		for {
			select {
			case msg1 := <-ch1:
				fmt.Println(msg1)
			case msg2 := <-ch2:
				fmt.Println(msg2)
			}
		}
	}()
	fmt.Println("Done")
}
