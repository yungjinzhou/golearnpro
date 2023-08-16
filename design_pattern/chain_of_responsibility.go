// 责任链模式
/*
责任链模式是一种行为型设计模式，旨在将请求的发送者和接收者解耦，并将其组织成一条链。每个对象都有机会处理请求，或者将其传递给链中的下一个对象，直到请求被处理或者到达链的末尾。

在责任链模式中，通常有以下角色：

抽象处理器（Handler）：定义了处理请求的接口，并且持有下一个处理器的引用。可以是抽象类或接口。
具体处理器（Concrete Handler）：实现了抽象处理器接口，负责处理请求。如果自己无法处理，则将请求转发给下一个处理器。
责任链模式的主要思想是请求沿着链传递，每个处理器根据自己的判断是否可以处理请求，如果能够处理，则直接处理，否则将请求传递给下一个处理器。这样，请求的发送者和接收者之间就解耦了，发送者无需知道请求会由谁处理，而处理者也无需知道请求的具体来源。

责任链模式的主要优点包括：

解耦性：发送者和接收者之间解耦，发送者无需知道请求会被谁处理，处理者也无需知道请求的具体来源。
增强灵活性：可以通过动态改变链的顺序或增加新的处理器来灵活地处理请求。
可扩展性：可以方便地增加新的处理器来处理新的请求类型，符合开闭原则。
然而，责任链模式也存在一些限制：

请求可能无法得到处理：如果整个链上的所有处理器都无法处理请求，那么请求可能会被忽略或被丢弃。
对请求的处理顺序敏感：责任链中处理器的顺序可能对请求的最终结果产生影响，需要谨慎配置处理器的顺序。
总结来说，责任链模式通过将请求发送者和接收者解耦，将其组织成一条链，并且每个处理器都有机会处理请求或将其传递给下一个处理器。它提供了一种灵活、可扩展的方式来处理请求，但也需要注意处理器顺序和请求无法得到处理的问题。责任链模式常用于请求的处理、事件驱动系统等场景。
*/

package main

import "fmt"

type Handle interface {
	handleLeave(day int)
}

type GeneralManager struct {
}

type DepartmentManager struct {
	next Handle
}

type ProjectDirector struct {
	next Handle
}

func (g *GeneralManager) handleLeave(day int) {
	if day <= 10 {
		fmt.Printf("总经理批准，假期 %d 天\n", day)
	} else {
		fmt.Printf("总经理权限不足，不能通过")
	}
}

func (d *DepartmentManager) handleLeave(day int) {
	if day <= 5 {
		fmt.Printf("部门经理批准，假期 %d 天\n", day)
	} else {
		fmt.Printf("部门经理权限不足\n")
		d.next.handleLeave(day)
	}
}

func (p *ProjectDirector) handleLeave(day int) {
	if day <= 3 {
		fmt.Printf("项目经理批准，假期 %d 天\n", day)
	} else {
		fmt.Printf("项目经理权限不足\n")
		p.next.handleLeave(day)
	}
}

func main() {
	day := 4
	dir := &ProjectDirector{next: &DepartmentManager{next: &GeneralManager{}}}
	dir.handleLeave(day)
}
