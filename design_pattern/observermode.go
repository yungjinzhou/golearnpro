package main

import "fmt"

// Observer 观察者接口
type Observer interface {
	Update(message string)
}

// Subject 主题（被观察者）对象
type Subject struct {
	observers []Observer
}

// RegisterObserver 注册观察者
func (s *Subject) RegisterObserver(observer Observer) {
	s.observers = append(s.observers, observer)
}

// RemoveObserver  移除观察者
func (s *Subject) RemoveObserver(observer Observer) {
	for i, obs := range s.observers {
		if obs == observer {
			s.observers = append(s.observers[:i], s.observers[i+1:]...)
			break
		}
	}
}

// NotifyObserver 通知观察者
func (s *Subject) NotifyObserver(message string) {
	for _, obs := range s.observers {
		obs.Update(message)
	}
}

// ConcreteObserver 具体观察者类型
type ConcreteObserver struct {
	name string
}

func (o *ConcreteObserver) Update(message string) {
	fmt.Printf("[%s] 收到消息：%s\n", o.name, message)
}

func main() {
	// 创建主体对象
	subject := &Subject{}

	// 创建观察者对象
	obs1 := &ConcreteObserver{name: "观察者1"}
	obs2 := &ConcreteObserver{name: "观察者2"}
	obs3 := &ConcreteObserver{name: "观察者3"}

	// 注册观察者
	subject.RegisterObserver(obs1)
	subject.RegisterObserver(obs2)
	subject.RegisterObserver(obs3)

	// 发送消息
	subject.NotifyObserver("测试")

	//移除观察者
	subject.RemoveObserver(obs1)

	// 发送消息
	subject.NotifyObserver("测试2")
}


//
//
//
//
//
//
//package main
//
//import "fmt"
//
//// Subscriber  订阅者
//type Subscriber interface {
//	Update(message string)
//}
//
//// Publisher 发布者
//type Publisher struct {
//	Subscribers []Subscriber
//}
//
//func (p *Publisher) RegistrySubscriber(subscriber Subscriber) {
//	p.Subscribers = append(p.Subscribers, subscriber)
//}
//
//func (p *Publisher) RemoveSubscriber(subscriber Subscriber) {
//	for i, sub := range p.Subscribers {
//		if sub == subscriber {
//			p.Subscribers = append(p.Subscribers[:i], p.Subscribers[i+1:]...)
//			break
//		}
//	}
//}
//
//func (p *Publisher) NotifySubscriber(message string) {
//	for _, subscriber := range p.Subscribers {
//		subscriber.Update(message)
//	}
//}
//
//type ConcreteSubscriber struct {
//	name string
//}
//
//func (sub *ConcreteSubscriber) Update(message string) {
//	fmt.Printf("订阅者：%s 收到消息%s\n", sub.name, message)
//}
//
//func main() {
//	publisher := &Publisher{}
//
//	sub1 := &ConcreteSubscriber{name: "订阅者1"}
//	sub2 := &ConcreteSubscriber{name: "订阅者2"}
//
//	publisher.RegistrySubscriber(sub1)
//	publisher.RegistrySubscriber(sub2)
//
//	publisher.NotifySubscriber("测试发布")
//
//	publisher.RemoveSubscriber(sub1)
//
//	publisher.NotifySubscriber("测试发布2")
//
//}
