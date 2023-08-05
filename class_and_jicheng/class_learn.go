package main

import "fmt"

/*

class Person():
   name = None
   xxxxx

*/

// struct类似python的class

type Person struct {
	name   string
	age    int
	gender string
	score  float64
}

// 在类外绑定方法

func (p *Person) Eat() {
	fmt.Sprintln("Person is eating")
	fmt.Sprintln(p.name + "is eating")
	p.name = "Duck"

}

func (p Person) Eat2() {
	fmt.Sprintln("Person is eating")
	fmt.Sprintln(p.name + "is eating")
	p.name = "duke"
}

func main() {
	lily := Person{
		name:   "lily",
		age:    30,
		gender: "女",
		score:  10,
	}
	lily2 := lily
	fmt.Println("Eat, 使用p  *Persion，使用指针....")
	fmt.Println("修改前，lily", lily) //lily
	lily.Eat()
	fmt.Println("修改后lily", lily) // duke

	fmt.Println("Eat2, 使用p Persion，不使用指针....")
	fmt.Println("修改前， lily2....", lily2) // lily
	lily2.Eat2()
	fmt.Println("修改后，lily2....", lily2) // lily

}
