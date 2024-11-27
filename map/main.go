package main

import (
	"fmt"
	"strings"
)

func main() {
	m3 := make(map[int]string)
	m3[1] = "a"
	fmt.Println(m3)
	fmt.Println(len(m3))

	m4 := make(map[int]string, 4)
	m4[1] = "a"
	m4[1] = "b"
	fmt.Println(m4)
	fmt.Println(len(m4))

	fmt.Println("=======")

	var mm1 map[int]string = map[int]string{1: "a"}
	mm1[200] = "b"
	fmt.Println(mm1)

	mm2 := map[int]string{1: "a"}
	mm2[200] = "b"
	fmt.Println(mm2)

	mm3 := make(map[int]string)
	mm3[100] = "a"
	mm3[200] = "b"
	mm3[300] = "c"
	mm3[400] = "d"
	if _, is := mm3[200]; is == true {
		fmt.Println("存在")
	} else {
		fmt.Println("不存在")
	}
	delete(mm3, 100)
	fmt.Println(mm3) // map[200:b 300:c 400:d]
	testDel(mm3, 200)
	fmt.Println(mm3) // map[300:c 400:d]

	fmt.Println("==stat==")
	statWord()
}

// 测试map作为参数传递是值传递还是引用传递
func testDel(mm map[int]string, key int) {
	delete(mm, key)
}

// 作业1: 统计字符串中的单词出现个数
func statWord() {
	str := "a b c d a d c e d"
	sliceStr := strings.Split(str, " ")
	mm := make(map[string]int)

	for _, v := range sliceStr {
		if vv, ok := mm[v]; ok {
			mm[v] = vv + 1
		} else {
			mm[v] = 1
		}
	}
	fmt.Println(mm)
}
