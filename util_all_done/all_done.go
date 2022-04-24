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


func AllResponse() string {
	numOfRunner := 10
	ch := make(chan string, numOfRunner)
	for i := 0;i < numOfRunner; i++{
		go func(i int) {
			ret := runTask(i)
			ch <- ret
		}(i)
	}
	finalRet := ""
	for j:=0;j < numOfRunner;j++{
		finalRet += <-ch + "\n"
	}
	return finalRet
}


func main() {

	fmt.Println("Before:", runtime.NumGoroutine())
	ret := AllResponse()
	fmt.Println(ret)
	time.Sleep(time.Second * 1)
	fmt.Println("After:", runtime.NumGoroutine())
}
