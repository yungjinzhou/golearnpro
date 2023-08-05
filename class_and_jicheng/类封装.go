package main

import "fmt"

type Human struct {
	name   string
	age    int
	gender string
}

func (h *Human) Eat() {
	fmt.Println("this is  : ,", h.name)
}

// 类的嵌套

type Student struct {
	hum    Human //属于类的嵌套
	school string
	score  float64
}

//类的继承

type Teacher struct {
	subject string
	Human   // 直接写类，没有字段名字
}

func main() {
	s1 := Student{
		hum: Human{
			name:   "lily",
			age:    18,
			gender: "女",
		},
		school: "sss",
		score:  10,
	}
	fmt.Println(s1.hum.name)
	fmt.Println(s1.school)

	t1 := Teacher{}
	t1.subject = "数学"
	t1.age = 80
	t1.gender = "女"
	t1.name = "闫老师"
	fmt.Println(t1.subject)
	// 子类默认会创建和类同名字段，为了在子类中能操作父类
	fmt.Println(t1.name)
	fmt.Println(t1.Human.name)
}
