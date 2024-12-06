package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	t3()
}

func t3() {
	ctx := context.WithValue(context.Background(), "trunk_id", "12346")

	go func(ctx context.Context) {
		v := ctx.Value("trunk_id")
		fmt.Println(v)
	}(ctx)

	time.Sleep(time.Second * 3)
	fmt.Println("end")
}

func t2() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	go func(ctx context.Context) {
		v := <-ctx.Done()
		fmt.Println(v)
	}(ctx)

	for i := 0; i < 10; i++ {
		fmt.Println(i)
		time.Sleep(1 * time.Second)
	}
}

func t1() {
	ctx, cancel := context.WithCancel(context.Background())

	go func(_ context.Context) {
		for i := 0; i < 10; i++ {
			time.Sleep(1 * time.Second)
			fmt.Println(i)
		}
	}(ctx)

	time.Sleep(5 * time.Second)
	cancel()
	time.Sleep(1 * time.Second)
}
