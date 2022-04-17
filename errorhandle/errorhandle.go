package main

import (
	"errors"
	"fmt"
)

func GetFibonacci(n int) ([]int, error) {

	if n <0 || n > 100{
		return nil, errors.New("n should be in [0,100]")
	}

	fibList := []int{1, 1}

	for i := 2;i<n;i++{
		fibList = append(fibList, fibList[i-2]+fibList[i-1])
	}
	return fibList, nil
}


func main(){
	ret,er := GetFibonacci(-10)
	fmt.Println(ret, er)
}



//func GetFibonacci(n int) []int {
//	fibList := []int{1, 1}
//
//	for i := 2;i<n;i++{
//		fibList = append(fibList, fibList[i-2]+fibList[i-1])
//	}
//	return fibList
//}
//
//
//func main(){
//	ret := GetFibonacci(10)
//	fmt.Println(ret)
//}
