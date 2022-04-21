package main

import (
	"fmt"
	"runtime"
	"time"
)

func runTask(id int) string {
	time.Sleep(10 * time.Millisecond)
	return fmt.Sprintf("The result is from %d", id)
}


func FirstResponse() string {
	numOfRunner := 10
	ch := make(chan string, numOfRunner)
	for i := 0;i < numOfRunner; i++{
		go func(i int) {
			ret := runTask(i)
			ch <- ret
		}(i)
	}
	return <-ch
}


func main() {

	fmt.Println("Before:", runtime.NumGoroutine())
	ret := FirstResponse()
	fmt.Println(ret)
	time.Sleep(time.Second * 1)
	fmt.Println("After:", runtime.NumGoroutine())
}



// first resposne
//func FirstResponse() string {
//	numOfRunner := 10
//	ch := make(chan string)
//	for i := 0;i < numOfRunner; i++{
//		go func(i int) {
//			ret := runTask(i)
//			ch <- ret
//		}(i)
//	}
//	return <-ch
//}
//
//
//func main() {
// test first response
//	ret := FirstResponse()
//	fmt.Println(ret)
//}