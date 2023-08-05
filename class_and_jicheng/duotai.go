package main

import "fmt"

//定义一个类型，类型是interface，接口中的函数是虚函数，不能有实现

type IAttack interface {
	Attack()
}

type HumanLowLevel struct {
	name  string
	level int
}

type HumanHighLevel struct {
	name  string
	level int
}

func (a *HumanLowLevel) Attack() {
	fmt.Println("我是：", a.name, "我的等级是：", a.level)
}

func (a *HumanHighLevel) Attack() {
	fmt.Println("我是：", a.name, "我的等级是：", a.level)
}

// 传入不同的对象，调用同样的方法，实现多态

func DoAttack(a IAttack) {
	a.Attack()

}

func main() {
	lowLevel := HumanLowLevel{
		name:  "David",
		level: 1,
	}
	lowLevel.Attack()

	var player IAttack // 定义一个接口变量

	highLevel := HumanHighLevel{
		name:  "lily",
		level: 10,
	}
	highLevel.Attack()

	player = &lowLevel
	player.Attack()

	player = &highLevel
	player.Attack()

	// 多态
	DoAttack(&lowLevel)
	DoAttack(&highLevel)

}
