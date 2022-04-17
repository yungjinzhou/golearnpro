package main

import "fmt"

type Pet struct {

}

func (p *Pet) Speak(){
	fmt.Print("....")
}

func (p *Pet) SpeakTo(host string) {
	p.Speak()
	fmt.Println(" ", host)
}

type Dog struct {
	p *Pet
}

func (d *Dog) Speak(){
	d.p.Speak()
	fmt.Print("...Wang!Wang!.")
}

func (d *Dog) SpeakTo(host string) {
	d.p.SpeakTo(host)
	//fmt.Println(" ", host)
}


func main() {
	dog := new(Dog)
	dog.SpeakTo("Yang")
	dog.Speak()
}
