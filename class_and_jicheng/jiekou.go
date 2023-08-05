package main

import "fmt"

func main() {
	var i, j, k interface{} //空接口
	names := []string{"duke", "lily"}
	i = names
	fmt.Println("爱表切片数组：", i)
	age := 20
	j = age
	fmt.Println("j代表数字", j)

	str := "字符串"
	k = str
	fmt.Println("k代表数字", k)

	//具体类型判断
	kvalue, ok := k.(int)
	if !ok {
		fmt.Println("不是int")
	} else {
		fmt.Println("是int", kvalue)
	}
	// 使用switch判断用户输入的不同类型，根据不同类型做相应的处理
	array := make([]interface{}, 3)
	array[0] = 1
	array[1] = "hello world"
	array[2] = true
	for _, v := range array {
		switch v.(type) {
		case string:
			fmt.Printf("类型是字符串，内容为 %s\n", v)
		case int:
			fmt.Printf("类型是字符串，内容为 %d\n", v)
		case bool:
			fmt.Printf("类型为bool，内容为 %v\n", v)
		default:
			fmt.Println("不合理的数据类型")
		}
	}

}
