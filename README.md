##### 一些简单的知识点
+ defer是在函数return, painc之前执行

+ defer的执行顺序是后来居上原则, 先定义后执行, 后定义限执行

+ make函数只能针对channel, map, slice, make会自动初始化类型然后返回的是引用类型

+ new返回指针类型, 返回一个地址, 并且不会初始化