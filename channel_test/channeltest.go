package main

import (
	"fmt"
	"sync"
)

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
