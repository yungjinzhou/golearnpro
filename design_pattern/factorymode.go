package main

import "fmt"

type Payment interface {
	pay(money int)
}

type WechatPayment struct {
}

type Alipayment struct {
}

func (w *WechatPayment) pay(money int) {
	fmt.Printf("微信支付 %d\n", money)
}

func (w *Alipayment) pay(money int) {
	fmt.Printf("支付宝支付 %d\n", money)
}

func PaymentFactory(method string) Payment {
	switch method {
	case "ali":
		return &Alipayment{}
	case "wechat":
		return &WechatPayment{}
	default:
		return nil
	}
}

func main() {
	alipay := PaymentFactory("ali")
	alipay.pay(100)

}
