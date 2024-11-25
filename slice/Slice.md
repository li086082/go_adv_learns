#### 切片解决的问题

1. 数组的容量固定, 不容自动扩展
2. 数组作为参数的时候会进行值传递给形参

##### 切片的本质

-   不是数组的指针, 是一种数据结构, 用来操作数组内部元素

##### 切片的使用

-   定义 : 名称[起始下标, 结束下标, 容量]
    -   容量 = 最大长度 - 起始下标
    -   不写容量则等于元素组容量
    -   长度 = 结束下标 - 起始下标

##### 数组和切片的区别

-   创建数组时，[]内需要必须写长度或者是...进行长度推导
-   创建切片时，[]内为空

##### 如何判断数组还是切片

:::info
fmt.Println(reflect.Typeof(arr).kind())

:::

##### 切片的创建

-   自动推导 s := []int{1,2}
-   s := make([]int, 长度, 容量)
-   s := make([]int, 长度) // 创建切片时没有指定容量, 容量等于长度

##### 切片做函数参数

-   地址传递
-   append 函数会生成一个新 slice, 所以被调用函数中进行 append 后不会影响调用函数切片

删除切片中指定元素

```go
func t4(s []int, i int, num int) []int {
    return append(s[:i], s[i+num:]...)
}
```

##### 深拷贝

```go
func t5(){
	// dst为空
	s1 := "hello"
	s2 := make([]byte, len(s1))
	aaa := copy(s2, s1)
	fmt.Println(aaa)
	fmt.Println(string(s2))

	// dst不为空, 会对同下标元素overlap
	ss1 := []int{1, 2, 3}
	ss2 := []int{4, 5, 6}
	ss3 := copy(ss1, ss2)
	ss2[0] = 100
	fmt.Println(ss3)
	fmt.Println(ss1)
	fmt.Println(ss2)
}
```
