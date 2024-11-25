package main

import (
	"fmt"
)

func main(){
	arr := [...]int{1, 2, 3, 4, 5, 6, 7}
	
	s := arr[:5]

	fmt.Println(s)
	fmt.Println(len(s))
	fmt.Println(cap(s))

	fmt.Println("=====")
	t1()

	fmt.Println("=====")
	t2()

	fmt.Println("=====")
	t3(s)
	fmt.Println(s)

	fmt.Println("=====")
	a := t4(s, 2, 2)
	fmt.Println(a)

	fmt.Println("=====t5")
	t5()
}

func t1(){
	arr := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	s1 := arr[2:5]
	fmt.Println(cap(s1))
	fmt.Println(len(s1))
	fmt.Println(s1)

	s2 := s1[1:3]
	fmt.Println(s2)
}

func t2(){
	
	s1 := []int{}
	fmt.Println(s1)

	s2 := make([]int, 2, 3)
	fmt.Println(s2)
	fmt.Println(len(s2))
	fmt.Println(cap(s2))

	s3 := make([]int, 2)
	fmt.Println(s3)
	fmt.Println(len(s3))
	fmt.Println(cap(s3))
}

func t3(a []int){
	a[0] = 100
	// a = append(a, 200) // no
}


func t4(s []int, i int, num int) []int {
	return append(s[:i], s[i+num:]...)
}

func t5(){
	// dst为空
	s1 := "hello"
	s2 := make([]byte, len(s1))
	aaa := copy(s2, s1)
	fmt.Println(aaa)
	fmt.Println(string(s2))

	// dst不为空, overlap
	ss1 := []int{1, 2, 3}
	ss2 := []int{4, 5, 6, 8}
	ss3 := copy(ss1, ss2)
	ss2[0] = 100
	fmt.Println(ss3)
	fmt.Println(ss1)
	fmt.Println(ss2)
	fmt.Println("====")

	a1 := []int{5, 6, 7, 8, 9}
	a2 := a1[2:]
	a3 := a1[2+1:]
	fmt.Println(a2) // [7, 8, 9]
	fmt.Println(a3) // [8, 9]
	copy(a2, a3)
	fmt.Println(a2) // [8, 9, 9]
}