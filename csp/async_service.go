package main

import (
	"fmt"
	"time"
)

func service() string {
	time.Sleep(time.Millisecond * 50)
	return "Done"
}


func otherTask() {
	fmt.Println("working on someting else")
	time.Sleep(time.Millisecond * 100)
	fmt.Println("Task is done.")
}

func ceshiService() {
	fmt.Println(service())
	otherTask()
}



func main() {
	ceshiService()
}
