package main

import (
	"fmt"
	"sync"
)

type Singleton struct {
	name string
}

var instance *Singleton
var once sync.Once

func GetInstance() *Singleton {
	once.Do(func() {
		instance = &Singleton{}
	})
	return instance
}

func (s *Singleton) SetName(name string) {
	s.name = name
}

func (s *Singleton) GetName() string {
	return s.name
}

func main() {
	s1 := GetInstance()
	s1.SetName("ins1")
	fmt.Println(s1.GetName())

	s2 := GetInstance()
	s2.SetName("ins2")
	fmt.Println(s2.GetName())

	fmt.Println(s1.GetName()) // 输出: test2 (与instance2共享同一实例)

}
