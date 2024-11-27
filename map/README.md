##### 别名
+ 字典
+ 映射

##### map特点
+ 唯一
+ 无序

##### 不能做key的类型, 引用类型不能为key, 因为这些类型的值不是固定的
+ slice
+ function
+ map

##### 思考
+ 利用"唯一"的特点, 可以进行去重

##### 创建
+  var m1 map[int]string // 空map
    + 注意: 给空map赋值会报错"panic: assignment to entry in nil map"
    + 要正确使用, 则需要在声明后机赋值

+  m2 := map[int]string{}
    + 默认长度为0, 可以赋值

+ m3 := make(map[int]string)
    + 等同于m2

+ m4 := make(map[int]string, 2)
    + 指定长度为2, 但是使用len(m4)还是为0

##### 初始化
+ var mm1 map[int]string = map[int]string{1: "ljx"}
    + 初始化的时候key值不能重复, 重复会报编译错误


+ mm2 := map[int]string{1: "hello"}
    + 初始化的时候key值不能重复, 重复会报编译错误

##### 遍历
+ for...range


##### 判断是否存在
```go
mm3 := make(map[int]string)
mm3[100] = "a"
mm3[200] = "b"
if _, ok := mm3[200]; is == ok {
    fmt.Println("存在")
} else {
    fmt.Println("不存在")
}
```
> mm3[100] 在等号左边是赋值, 接收1个值

> mm3[100] 在等号右边是取值, 返回2个值, 第一个是mm3[100]存储值, 第二个是100这个下标是否存在的bool值

##### 删除一个值
+ delete(mm3, 100)
    + 第一个参数是需要操作的map
    + 第二个参数是需要删除的key
    + 支持删除一个不存在的key,不会报错

##### map作为参数传递
+ map在作为函数参数传递时是地址传递

##### 作业1
```go
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
	fmt.Println(mm) // map[a:2 b:1 c:2 d:3 e:1]
}

```

##### 备注
+ map不支持cap函数