package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func run() {

}

type Window interface {
	start()
	repaint()
	stop()
	run()
}

type MyWindow struct {
	msg string
}

func (w *MyWindow) start() {
	fmt.Println("开始运行")
}

func (w *MyWindow) stop() {
	fmt.Println("结束运行")
}

func (w *MyWindow) repaint() {
	fmt.Println(w.msg)
}

func (w *MyWindow) run() {
	w.start()
	sigalCh := make(chan os.Signal, 1)
	signal.Notify(sigalCh, os.Interrupt, syscall.SIGTERM)
	for {
		select {
		case <-time.After(time.Second):
			w.repaint()
		case <-sigalCh:
			w.stop()
			return
		}
	}
}

func main() {
	ww := MyWindow{
		msg: "xxxx",
	}
	ww.run()
}
