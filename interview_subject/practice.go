package main

import (
	"fmt"
)


import "encoding/json"

// tag使用
type Stu struct {
	Name string `json:"stu_name"`
	ID   string `json:"stu_id"`
	Age  int    `json:"-"`
}



// 交换两个值
//func practicesubject() {
//	a, b := "A", "B"
//	a, b = b, a
//	fmt.Println(a, b) // B A
//}



//如何判断 map 中是否包含某个 key
//func practicesubject() {
//	shuzu := [...]float32{1000.0, 2.0, 3.4, 7.0, 50.0}
//	var aa = make(map[string]int)
//	if val, ok := aa["foo"]; ok {
//		fmt.Println(val)
//		fmt.Println(ok)
//	}
//	fmt.Println(shuzu)
//}



// 字符串测试
//func practicesubject() {
//	var str strings.Builder
//	for i := 0; i < 1000; i++ {
//		str.WriteString("a")
//	}
//	fmt.Println(str.String())
//}


// 指针测试
//func practicesubject() {
//	var a = 10
//	var p *int = &a
//	fmt.Println(p)
//}



func main(){
	buf, _ := json.Marshal(Stu{"Tom", "t001", 18})
	fmt.Printf("%s\n", buf)
}