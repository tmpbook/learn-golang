package main

import (
	"context"
	"fmt"
	"time"
)

func C(ctx context.Context) error {
	fmt.Println(ctx.Value("key3"))
	select {
	// 结束时候做点什么 ...
	case <-ctx.Done():
		return ctx.Err()
	default:
		// 没有结束 ... 执行 ...
		return nil
	}
}

func B(ctx context.Context) int {
	fmt.Println(ctx.Value("key0"))
	fmt.Println(ctx.Value("key1"))
	ctx = context.WithValue(ctx, "key3", "value3")
	go fmt.Println(C(ctx))
	select {
	// 结束时候做点什么 ...
	case <-ctx.Done():
		return -2
	default:
		// 没有结束 ... 执行 ...
		return 0
	}
}

func A(ctx context.Context) int {
	ctx = context.WithValue(ctx, "key0", "value0")
	ctx = context.WithValue(ctx, "key1", "value1")
	go fmt.Println(B(ctx))
	select {
	// 结束时候做点什么 ...
	case <-ctx.Done():
		return -1
	default:
		// 没有结束 ... 执行 ...
		return 0
	}
}
func main() {
	// 定时取消
	// {
	// 	timeout := 3 * time.Second
	// 	ctx, _ := context.WithTimeout(context.Background(), timeout)
	// 	fmt.Println(Add(ctx))
	// }

	// 手动取消
	{
		ctx, cancel := context.WithCancel(context.Background())
		go func() {
			time.Sleep(2 * time.Second)
			// 主动取消
			cancel()
		}()
		fmt.Println(A(ctx))
	}
	select {}
}
