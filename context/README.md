##### 如何创建一个Context
+ Context.Background()
+ Context.TODO()

##### Context四个功能方法
+ WithCancel
+ WithDeadline
+ WithTimeout
+ WithValue

##### WithCancel
+ 在使用WithCancel时会返回一个cancel方法
+ 在主goroutine中调用cancel方法时
+ 子goroutine的ctx.Done通道会接收到一个空结构的信息

##### WithDeadline - WithTimeout
+ 这两个用法是一样的, 本质上也是一样timeout内部调用的也是deadline
+ 在创建使用WithTimeout创建context时, 可以设置一个时间参数
+ 子goroutine在时间参数到达后ctx.Done通道会接收到一个空结构的信息

#### WithValue
+ context传值, 建议传递基础数据类型
+ 子goroutine只能取值, 不能改值