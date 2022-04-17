package main

import (
	"fmt"
	"time"
)

// 未看到效果，还是顺序执行

func TestGroutine(){
	for i := 0; i < 10; i++ {
		go func(i int) {
			fmt.Println(i)
		}(i)

		//go func() {
		//	fmt.Println(i)
		//}()
		time.Sleep(time.Millisecond* 50)
	}
 }


func main() {
	TestGroutine()

}
