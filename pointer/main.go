package main

import "fmt"

func swap(a, b int) (int, int) {
	return b, a
}

func swap2(a, b *int) {
	*a, *b = *b, *a
}

func main() {

	a, b := 10, 20
	a1, b1 := swap(a, b)
	fmt.Println(a, b)
	fmt.Println(a1, b1)

	swap2(&a, &b)
	fmt.Println(a, b)

	fmt.Println("===开始烧脑了===")

	aa1 := 20
	var aa2 *int = &aa1  // 一级指针
	var aa3 **int = &aa2 // 二级指针
	fmt.Println(aa1)
	fmt.Println(aa2)
	fmt.Println(aa3)
	fmt.Println(*aa2)
	fmt.Println(**aa3)
	fmt.Println("===aa2重新赋值===")
	*aa2 = 30
	fmt.Println(aa1)
	fmt.Println(*aa2)
	fmt.Println(*aa3)
	fmt.Println(**aa3)

	fmt.Println("===aa3重新赋值===")
	**aa3 = 50
	fmt.Println(aa1)
	fmt.Println(*aa2)
	fmt.Println(**aa3)
}
