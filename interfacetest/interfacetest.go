package main

import (
	"fmt"
	"time"
)

// 自定义类型结合接口用例测试


type IntConvertor func(op int) int


func timeSpent(inner IntConvertor) IntConvertor {
		return func(n int) int {
			start := time.Now()
			ret := inner(n)
			fmt.Println("time spent:", time.Since(start).Seconds())
			return ret
		}

//func timeSpent(inner func(op int) int) func (op int) int {
//	return func(n int) int {
//		start := time.Now()
//		ret := inner(n)
//		fmt.Println("time spent:", time.Since(start).Seconds())
//		return ret
//	}
}

func slowFun(op int) int{
	fmt.Println("this is slow fun function")
	time.Sleep(time.Second*1)
	return op
}


func main() {
	//spent := timeSpent(slowFun)
	//fmt.Println(spent(10))
	spent := timeSpent(slowFun)(10)
	fmt.Println(spent)
}





// 接口用例测试

//type Programmer interface {
//	WriteHelloWorld() string
//}
//
//type GoProgrammer struct {
//
//}
//
//func (g *GoProgrammer) WriteHelloWorld() string {
//	return "fmt.Println(\"Hello World!\")"
//}
//
//func main(){
//	var p Programmer
//	p = new(GoProgrammer)
//	ret := p.WriteHelloWorld()
//	fmt.Println(ret)
//}




// 接口实现

//type GoProgrammer struct {
//}
//
//func (p *GoProgrammer) WriteHelloWorld() Code {
//	return "fmt.Println(\"Hello World\")"
//}



// 接口定义
//type Programmer interface(){
//	WriteHelloWorld() Code
//}







