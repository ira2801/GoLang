package main

import (
	"fmt"
)

func main() {

	generates := make(chan int)
	filters := make(chan int)
	squares := make(chan int)
	sums := make(chan int)

	go func() {
		for i := 1; i <= 100; i++ {
			generates <- i
		}
		close(generates)
	}()
	go func() {
		for generate := range generates {
			if generate%2 == 0 {
				filters <- generate
			}
		}
		close(filters)
	}()
	go func() {
		for filter := range filters {
			squares <- filter * filter
		}
		close(squares)
	}()
	go func() {
		total := 0
		for {
			sum, ok := <-squares
			if !ok {
				break
			}
			total += sum
		}
		sums <- total
		close(sums)
	}()
	fmt.Println("Фінальна сума:", <-sums)
}
