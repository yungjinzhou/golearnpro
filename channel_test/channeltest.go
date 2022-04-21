package main

import (
	"fmt"
	"sync"
)

//向关闭的channel发送数据，会导致,panic
//v, ok <-ch; ok为bool值，true表示正常接受，false表示通道关闭
//所有的channel接收者都会在channel关闭时，立刻从阻塞等待中返回且.上述ok值为false。这个广播机制常被利用，进行向多个订阅者同时发送信号。
//如:退出信号。

// v, ok <- ch;  ok为bool值，true 为正常接收，false代表channel关闭

//func dataProducer(ch chan int, wg *sync.WaitGroup) {
//	go func() {
//		for i := 0;i <10;i++ {
//			ch <- i
//		}
//		wg.Done()
//	}()
//}
//
//
//
//
//func dataReceiver(ch chan int, wg *sync.WaitGroup) {
//	go func() {
//		for {
//			if data, ok := <-ch; ok {
//				fmt.Println(data)
//				} else{
//					break
//				}
//			}
//		}()
//}




// 指定个数
func dataProducer(ch chan int, wg *sync.WaitGroup) {
	go func() {
		for i := 0;i <10;i++ {
			ch <- i
		}
		wg.Done()
	}()
}


func dataReceiver(ch chan int, wg *sync.WaitGroup) {
	go func() {
		for  i := 0;i <10;i++{
				data := <-ch
				fmt.Println(data)
			}
			wg.Done()
}()
}


func main() {
	var wg sync.WaitGroup
	ch := make(chan int)
	wg.Add(1)
	dataProducer(ch, &wg)
	wg.Add(1)
	dataReceiver(ch, &wg)
	wg.Wait()
}
