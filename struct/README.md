##### 结构体
+ 结构体是一种数据类型

##### 声明
+ 使用type关键字, Ex: type a struct{}


##### 初始化
+ 全部成员都初始化, 此时可以不写属性名称, 对其数量和类型即可
+ 指定成员初始化, 未指定的成员会默认成类型初始值

##### 结构体的比较
+ 结构体提交只能比较==,!=
+ go语言内部会对比每个成员

##### 结构体之间赋值
+ 相同数据类型才能进行相互赋值, 也就是必须要同一个结构体
+ 结构体之间的赋值是值传递

##### 参数传递
+ 默认是值传递, 想当于实参拷贝一份赋值给形参, 推荐使用引用传递

##### 结构体嵌套(匿名嵌套和非匿名嵌套)
+ 一个具名结构体可以当做另一个的匿名属性(匿名嵌套)和非匿名属性(非匿名嵌套)来使用
+ 匿名嵌套的属性名默认和类型一致的, 在初始化的时候需要注意这个点
+ 在使用非匿名嵌套的时候,必须要按层级依次调用
+ 匿名嵌套在GO中会自动提升, 在使用的时候可以直接通过第一层结构使用, 当然也可依次调用
+ 匿名嵌套在遇到有重复属性的时候, 可以通过调用二层结构调用到这个值

##### 指针结构体的使用
+ 在使用var声明指针结构的时候需要将类型声明为*{struct}
+ 在使用方面结构体指针和普通结构体的用法是一致的
+ 结构体的指针地址就是结构体第一个元素的地址
```go

```