package main

import "fmt"

type NewPayment interface {
	pay(money int)
}

type Alipay struct {
}

func (a *Alipay) pay(money int) {
	fmt.Println("支付宝支付 %d 元", money)
}

type Wechatpay struct {
}

func (w *Wechatpay) pay(money int) {
	fmt.Println("微信支付 %d 元", money)
}

type BankPay struct {
}

func (b *BankPay) cost(money int) {
	fmt.Println("银联支付 %d 元", money)
}

type ApplePay struct {
}

func (ap *ApplePay) cost(money int) {
	fmt.Println("苹果支付 %d 元", money)
}

type NewBankPay struct {
	BankPay
}

// 第一种实现方式：类适配器，兼容payment和bankpay，可以实现两个类调用
//func (nb *NewBankPay) pay(money int) {
//	nb.cost(money)
//}
//
//func main() {
//	p1 := &NewBankPay{}
//	p1.pay(100)
//	p1.cost(100)
//}

// 第二种实现方式
type PaymentAdapter struct {
	payment interface{ cost(int) }
	//bankpay BankPay
}

func (pm *PaymentAdapter) pay(money int) {
	pm.payment.cost(money)
}

func main() {
	p := &PaymentAdapter{&ApplePay{}}
	p.pay(100)
	b := &PaymentAdapter{&BankPay{}}
	b.pay(100)
}
