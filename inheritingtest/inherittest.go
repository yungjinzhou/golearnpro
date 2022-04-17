package main

import (
	"fmt"
)

type Code string

type Programmer interface {
	WriteHelloWorld() Code
}

type GoProgrammer struct {
}

func (p *GoProgrammer) WriteHelloWorld() Code {
	return "fmt.Println(\"Hello world\")"
}

type  JavaProgrammer struct {

}

func (p *JavaProgrammer) WriteHelloWorld() Code {
	return "System.out.Println(\"Hello world\")"
}


func writeFirstProgram(p Programmer){
	fmt.Printf("%T %v\n", p, p.WriteHelloWorld())
}

func TestPolymorphism(){
	goProg := new(GoProgrammer)
	javaProg := new(JavaProgrammer)
	writeFirstProgram(goProg)
	writeFirstProgram(javaProg)
}


func main() {
	TestPolymorphism()
}

