package main

import (
	"fmt"
	"time"
)

func service() string {
	time.Sleep(time.Millisecond * 500)
	return "Done"
}


func AsyncService() chan string {
	retCh := make(chan string, 1)
	go func() {
		ret := service()
		fmt.Println("returned result.")
		retCh <- ret
		fmt.Println("Service exited.")
	}()
	return retCh
}



//使用select方式
func main () {
	select {
	case ret := <-AsyncService():
		fmt.Println(ret)
	case <- time.After(time.Millisecond * 100):
		fmt.Println("time out-----------")
	}
}


// 接收
//func main () {
//	ret := <- AsyncService()
//	fmt.Println(ret)
//}
