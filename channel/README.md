##### channel
+ 管道

##### 作用
+ 用户不同goroutine之间的相互通信

##### channel类型
+ 有缓冲区channel
+ 无缓冲区channel

##### 无缓冲区特点
+ 无缓冲区channel, 是同步的, 读写都会阻塞

##### 有缓冲区特点
+ 有缓冲区channel, 是异步的, 只有缓存区满时发送者会阻塞, 缓存区为空, 接收方会阻塞

##### 关闭channel
+ 发送协程和接收协程都能执行关闭channel的操作
+ 不能对已经关闭的channel再次使用close(), 会提示all goroutines are asleep - deadlock
+ 手动关闭(不是defer)代表没有消息再发送到channel上, 但是依然可以取值


#### 代码说明问题1
```go
ch := make(chan int)
defer close(ch)
ch <- 1
fmt.Println(<-ch)
```
+ 这段代码会报all goroutines are asleep - deadlock
+ 原因是在同一个goroutine中, 无缓存channel是同步的, 在写入和读取的时候会阻塞
+ 看第二行代码这个写入操作, 执行到了这一行后会一直阻塞在这里
+ 这样就执行不到下一行的读取操作, 形成了死锁
+ 解决办法有两种
```go
ch := make(chan int, 1)
defer close(ch)
ch <- 1
fmt.Println(<-ch)
```
+ 在不同的goroutine中分别进行读/写
```go
ch := make(chan int, 1)
defer close(ch)
go func() {
    ch <- 2
}()
fmt.Println(<-ch)
```

#### 代码说明问题2
```go
ch := make(chan int)
defer close(ch)

go func() {
    for i := 0; i < 4; i++ {
        ch <- i
    }
    for i := 0; i < 4; i++ {
        ch <- i + 4
    }
}()

for i := 0; i < 20; i++ {
    fmt.Println(<-ch)
}
```
+ 这段代码会引发deadlock, 原因是无缓冲channel是同步的, 
+ 子goroutine中写入了8次数据, 主goroutine中读取了20次, 
+ 在主goroutine读取第9次的时候就触发了deadlock, 它一直在读, 但是没有人写入
+ 还需要注意的一个点是, 这里的defer并不会执行
+ 解决办法有3中
```go
// 对其次数, 防止因为读写的死锁
ch := make(chan int)
defer close(ch)

go func() {
    for i := 0; i < 4; i++ {
        ch <- i
    }

    for i := 0; i < 4; i++ {
        ch <- i + 4
    }
}()

for i := 0; i < 8; i++ {
    fmt.Println(<-ch)
}
```

```go
// 在协程中写完后关闭channel
// 在手动关闭channel后, 是可以继续读数据的, 会读取到零值, 但是不能写数据, 会painc
ch := make(chan int)

go func() {
    defer close(ch)
    for i := 0; i < 4; i++ {
        ch <- i
    }

    for i := 0; i < 4; i++ {
        ch <- i + 4
    }
}()

for i := 0; i < 20; i++ {
    fmt.Println(<-ch)
}
```

#### 代码说明问题3
```go
ch := make(chan int, 2)
defer close(ch)
go func() {
    ch <- 11
    ch <- 22
}()
for v1 := range ch {
    fmt.Println(v1)
}
```
+ 这段代码会出现deadlock, 主要原因出在for...range...中
+ for...range...读取一个通道时会阻塞协程一直读取数据, 直到通道被close关闭, 在接收到close信息时, for才会退出
+ 这段代码中defer在for后面执行, 又因为for阻塞了, defer close其实并没有被执行

+ 解决办法有
```go
ch := make(chan int, 2)
go func() {
    defer close(ch)
    ch <- 11
    ch <- 22
}()
for v1 := range ch {
    fmt.Println(v1)
}
```
```go
ch := make(chan int, 2)
defer close(ch)
go func() {
    ch <- 11
    ch <- 22
}()
for i := 0; i < 2; i++ {
    fmt.Println(<-ch)
}
```