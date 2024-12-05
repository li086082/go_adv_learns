package main

import (
	"fmt"
	"time"
)

func main() {
	t8()
}

// 控制gotoutine请求时间
func t8() {
	out := make(chan bool, 1)
	defer close(out)
	go func() {
		time.Sleep(time.Second * 2)
		out <- true
	}()

	done := make(chan bool, 1)
	defer close(done)

	ch := make(chan int)
	defer close(ch)

	go func() {
		select {
		case <-done:
			return
		case <-time.After(time.Second * 3):
			ch <- 3
		}
	}()

	// 只要case有一个会满足, select就会运行完成
	select {
	case <-ch:
		fmt.Println("ch read data")
	case <-out:
		fmt.Println("time out")
		done <- true
	}
	time.Sleep(time.Second * 3)
	fmt.Println("run end")
}

func t7() {
	ch := make(chan int, 1)
	defer close(ch)
	ch <- 1
	// 当两个分支都会被满足时, 会随机执行一个分支
	select {
	case <-ch:
		fmt.Println(11)
	case <-ch:
		fmt.Println(22)
	default:
		fmt.Println("default")
	}
}

func t6() {
	ch := make(chan int, 1)
	defer close(ch)
	// 所有条件都不满足, 则走default
	// 如果没有default则会deadlock, 代码会在<-ch一直阻塞
	select {
	case <-ch:
		fmt.Println(11)
	default:
		fmt.Println("default")
	}
}

func t5() {
	ch := make(chan int, 1)
	defer close(ch)
	go func() {
		time.Sleep(time.Second * 3)
		ch <- 11
		ch <- 22
	}()
	// 这里会阻塞
	for i := 0; i < 2; i++ {
		fmt.Println(<-ch)
	}
	// 123会在上面循环打印结束后再打印
	fmt.Println(123)
}

func t4() {
	ch := make(chan int, 2)
	defer close(ch)
	go func() {
		ch <- 11
		ch <- 22
	}()
	for i := 0; i < 2; i++ {
		fmt.Println(<-ch)
	}
}

func t3() {
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
}

// 这段代码是正确的
func t2() {
	ch := make(chan int)
	close(ch)
	fmt.Println(<-ch) // 0
}

func t1() {
	ch := make(chan int)
	defer close(ch)

	go func() {
		ch <- 11
		ch <- 20
	}()

	for i := 0; i < 2; i++ {
		v1 := <-ch
		fmt.Println(v1)
	}
}
