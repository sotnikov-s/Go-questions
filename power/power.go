package main

import (
	"fmt"
	"time"
	"math/rand"
)

func main() {

	ch1 := make(chan int)
	closed := false
	num := 0

	go func() {
		for {
			if closed {
				break
			}
			num = rand.Intn(100)
			fmt.Println(num, "generated")
			ch1 <- num
			time.Sleep(1*time.Second)
		}
	}()

CheckLoop:
	for {
		select {
		case <-ch1:
			val := <-ch1
			if val > 90 {
				closed = true
				close(ch1)
				fmt.Println("CHANNEL CLOSED!")
				break CheckLoop
			}
			fmt.Println("Number read:", val)
		}
	}

	fmt.Println("\n\n-------------\nTHE END")
}
