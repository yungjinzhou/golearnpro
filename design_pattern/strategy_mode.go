package main

import "fmt"

type Strategy interface {
	execute(data string)
}

type FastStrategy struct {
}

func (f *FastStrategy) execute(data string) {
	fmt.Printf("用较快的策略处理 %s", data)
}

type SlowStrategy struct {
}

func (s *SlowStrategy) execute(data string) {
	fmt.Printf("用较慢的策略处理 %s", data)
}

type Context struct {
	data     string
	strategy Strategy
}

func (c *Context) setStrategy(strategy Strategy) {
	c.strategy = strategy
}

func (c *Context) doStrategy() {
	c.strategy.execute(c.data)
}

func main() {
	data := "xxx"
	s1 := &FastStrategy{}
	//s2 := SlowStrategy{}
	context := Context{data, s1}
	context.doStrategy()
}
