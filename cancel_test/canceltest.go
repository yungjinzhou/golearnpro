package main

import (
	"fmt"
	"time"
)

func isCancelled(cancelChan chan struct{}) bool {
	select {
	case <-cancelChan:
		return true
	default:
		return false
	}
}


func cancel1(cancelChan chan struct{}) {
	cancelChan <- struct{}{}
}


func cancel2(cancelChan chan struct{}) {
	close(cancelChan)
}

func ceshi_cancel() {
	cancelChan := make(chan struct{}, 0)
	for i := 0; i < 5; i++ {
		go func(i int, cancelCh chan struct{}) {
			for {
				if isCancelled(cancelCh) {
					break
				}
				time.Sleep(time.Millisecond * 5)
			}
			fmt.Println(i, "Done")
		}(i, cancelChan)
	}
	cancel1(cancelChan)
	cancel2(cancelChan)
	time.Sleep(time.Second * 1)
}

func main() {
	ceshi_cancel()
}