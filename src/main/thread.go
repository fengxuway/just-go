package main

import (
	"fmt"
)

func main() {
	c1, c2 := make(chan int), make(chan string)
	f := make(chan bool, 2)
	go func() {
		for {
			select {
			case v, ok := <-c1:
				if !ok {
					f <- true
					break
				}
				fmt.Println("c1: ", v)
			case v, ok := <-c2:
				if !ok {
					f <- true
					break
				}
				fmt.Println("c2: ", v)
			}
		}
	}()
	c1 <- 1
	c2 <- "dksl"
	c1 <- 3
	c1 <- 5
	c1 <- 6
	c1 <- 7
	c2 <- "dkslddddd"

	c2 <- "dkslddddd5"
	close(c1)
	for i:= 0; i< 2; i++{
		<- f
	}
}
