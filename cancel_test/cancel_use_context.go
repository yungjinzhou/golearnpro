package main

import (
	"context"
	"fmt"
"time"
)


// 根 Context: 通过 context.Background()创建
//子Context: context. WithCancel(parentContext)创建
//ctx, cancel := context.WithCancel(context. Background0)
//当前Context被取消时，基于他的子context都会被取消
//接收取消通知<-ctx.Done()


func isCancelled2(ctx context.Context) bool {
	select {
	case <-ctx.Done():
		return true
	default:
		return false
	}
}


func ceshiCancel2() {
	ctx, cancel := context.WithCancel(context.Background())
	for i := 0; i < 5; i++ {
		go func(i int, ctx context.Context) {
			for {
				if isCancelled2(ctx) {
					break
				}
				time.Sleep(time.Millisecond * 5)
			}
			fmt.Println(i, "Cancelled")
		}(i, ctx)
	}
	cancel()
	time.Sleep(time.Second * 1)
}

func main() {
	ceshiCancel2()
}
