package main

import (
	"fmt"
	"sync"
)


// waitgroup 机制，启动协程后加入等待组，所有协程运行完成后才会向下运行

func testCounterWaitGroup() {
	var mut sync.Mutex
	var wg sync.WaitGroup
	counter := 0
	for i :=0;i<5000;i++ {
		wg.Add(1)
		go func(){
			defer func () {
				mut.Unlock()
			}()
			mut.Lock()
			counter++
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Printf("couter = %d", counter)

}

func main(){
	testCounterWaitGroup()
}


//锁机制，线程安全，手动等待

//func testCounter() {
//	var mut sync.Mutex
//	counter := 0
//	for i :=0;i<5000;i++ {
//		go func(){
//			defer func () {
//				mut.Unlock()
//			}()
//			mut.Lock()
//			counter++
//		}()
//	}
//	time.Sleep(1 * time.Second)
//	fmt.Printf("couter = %d", counter)
//
//}
//
//func main(){
//	testCounter()
//}



// 没有锁机制
//func testCounter() {
//	counter := 0
//	for i :=0;i<5000;i++ {
//		go func(){
//			counter++
//		}()
//	}
//	time.Sleep(1 * time.Second)
//	fmt.Printf("couter = %d", counter)
//
//}
//
//func main(){
//	testCounter()
//}