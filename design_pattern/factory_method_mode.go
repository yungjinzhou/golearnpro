package main

import "fmt"

type PayMent interface {
	PayMethod(money int)
}

type PayMentFactory interface {
	CreatePayment() PayMent
}

type WechatPayMent struct {
}

type AliPayMent struct {
}

func (we *WechatPayMent) PayMethod(money int) {
	fmt.Printf("微信支付 %d\n", money)
}

func (ali *AliPayMent) PayMethod(money int) {
	fmt.Printf("支付宝支付 %d\n", money)
}

type AlipaymentFactory struct {
}

type WechatPaymentFactory struct {
}

func (Ali *AlipaymentFactory) CreatePayment() PayMent {
	return &AliPayMent{}
}

func (We *WechatPaymentFactory) CreatePayment() PayMent {
	return &WechatPayMent{}
}

func main() {
	alifactory := &AlipaymentFactory{}
	ali := alifactory.CreatePayment()
	ali.PayMethod(100)
}

//
//type Product interface {
//	Use(param string)
//}
//
//type Factory interface {
//	CreateProduct() Product
//}
//
//type ConcreteProductA struct{}
//
//func (p *ConcreteProductA) Use(param string) {
//	fmt.Printf("Using ConcreteProductA with parameter: %s\n", param)
//}
//
//type ConcreteProductB struct{}
//
//func (p *ConcreteProductB) Use(param string) {
//	fmt.Printf("Using ConcreteProductB with parameter: %s\n", param)
//}
//
//type ConcreteFactoryA struct{}
//
//func (f *ConcreteFactoryA) CreateProduct() Product {
//	return &ConcreteProductA{}
//}
//
//type ConcreteFactoryB struct{}
//
//func (f *ConcreteFactoryB) CreateProduct() Product {
//	return &ConcreteProductB{}
//}
//
//
//func main() {
//	factoryA := &ConcreteFactoryA{}
//	productA := factoryA.CreateProduct()
//	productA.Use("ParamA") // 输出：Using ConcreteProductA with parameter: ParamA
//
//	factoryB := &ConcreteFactoryB{}
//	productB := factoryB.CreateProduct()
//	productB.Use("ParamB") // 输出：Using ConcreteProductB with parameter: ParamB
//}
