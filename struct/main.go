package main

import (
	"fmt"
	"unsafe"
)

type A struct {
	name string
	age  int64
}

type BB struct {
	aa     A // 非匿名嵌套
	height int64
}

type CC struct {
	A      // 匿名嵌套
	height int64
}

type DD struct {
	A
	hobby []string
}

func main() {
	// start 初始化
	a := A{name: "a"}
	b := A{name: "a", age: 18}
	ab := A{"a", 20}
	fmt.Println(a, b, ab) // {a 0} {a 18} {a 20}

	var c A
	fmt.Println(c)
	c.name = "a"
	c.age = 18
	fmt.Println(c)
	// end start

	// start 结构体比较
	a1 := A{name: "a", age: 18}
	a2 := A{name: "a", age: 18}
	if a1 == a2 {
		fmt.Println("a1 = a2")
	} else {
		fmt.Println("a1 != 2")
	}
	// end 结构体比较

	// start 结构体之间的互相赋值
	b1 := A{name: "b", age: 20}
	var b2 A
	b2 = b1
	fmt.Printf("%p", &b1)
	fmt.Println()
	fmt.Printf("%p", &b2)
	fmt.Println()
	b1.name = "bb"
	fmt.Println(b1, b2)
	fmt.Println(unsafe.Sizeof(b1))
	fmt.Println(unsafe.Sizeof(b2))
	// end 结构体之间的互相赋值

	// start 结构非匿名嵌套
	c1 := A{name: "c", age: 22}
	c2 := BB{aa: c1, height: 180}
	fmt.Println(c2)
	c3 := BB{aa: A{name: "c3", age: 22}, height: 190}
	fmt.Println(c3.aa.name)
	fmt.Println(c3)
	// end 结构非匿名嵌套

	// start 结构匿名嵌套
	c5 := CC{A: A{name: "ljx", age: 2000}, height: 170}
	fmt.Println(c5.age)
	fmt.Println(c5.A.name)
	fmt.Println(c5.height)
	// end 结构匿名嵌套

	// start 结构体指针
	dd1 := &A{name: "dd", age: 23}
	fmt.Printf("%p", dd1)
	fmt.Println()

	var dd2 *A = &A{name: "dd", age: 23}
	fmt.Printf("%p", dd2)
	fmt.Println()
	fmt.Println(dd2.name)

	dd3 := new(A)
	dd3.name = "dd3"
	dd3.age = 20
	fmt.Printf("%p", dd3) // 结构体地址
	fmt.Println()
	fmt.Printf("%p", &dd3.name) // 结构体第一个属性地址
	fmt.Println()
	fmt.Printf("%p", &dd3.age)
	fmt.Println(dd3)
	// end

	// start 初始化带切片的结构体
	ddd1 := DD{A: A{name: "ddd1", age: 20}, hobby: []string{"eat"}}
	fmt.Println(ddd1)
	// end 初始化带切片的结构体
}
