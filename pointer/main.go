package main

import "fmt"

func swap(a, b int) (int, int){
	return b, a
}

func swap2(a, b *int){
	*a, *b = *b, *a
}

func main(){
	
	a, b := 10, 20
	a1, b1 := swap(a, b)
	fmt.Println(a, b)
	fmt.Println(a1, b1)

	swap2(&a, &b)
	fmt.Println(a, b)
}