package main

import "fmt"

func main() {

	ch := make(chan int)
	defer close(ch)

	arr := [...]string{"A", "B", "C", "D", "E"}

	go func() {
		ch <- 11
	}()

	v := <-ch

	fmt.Println(v)
}
