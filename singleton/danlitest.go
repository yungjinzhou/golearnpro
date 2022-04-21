package main

import (
	"fmt"
	"sync"
	"unsafe"
)

type Singleton struct {

}


var singleInstance *Singleton
var once sync.Once


func GetSingletonObj() *Singleton {
	once.Do(func() {
		fmt.Println("CREATE obj-------")
		singleInstance = new(Singleton)
	})
	return singleInstance
}


func main() {
	var wg sync.WaitGroup
	for i:=0; i<10;i++{
		wg.Add(1)
		go func() {
			obj := GetSingletonObj()
			//fmt.Println("%x\n", obj)
			fmt.Printf("%x\n", unsafe.Pointer(obj))
			wg.Done()
		}()
	}
	wg.Wait()
}


