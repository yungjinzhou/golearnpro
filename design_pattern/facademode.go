package main

import "fmt"

type CPU struct {
}

func (c *CPU) run() {
	fmt.Printf("cpu开始运行")
}

func (c *CPU) stop() {
	fmt.Printf("cpu停止运行")
}

type Memory struct {
}

func (m *Memory) run() {
	fmt.Printf("内存开始运行")
}

func (m *Memory) stop() {
	fmt.Printf("内存停止运行")
}

type Disk struct {
}

func (d *Disk) run() {
	fmt.Printf("磁盘开始运行")
}

func (d *Disk) stop() {
	fmt.Printf("磁盘停止运行")
}

type Compute struct {
	cpu    CPU
	memory Memory
	disk   Disk
}

func (c *Compute) run() {
	c.cpu.run()
	c.memory.run()
	c.disk.run()
}

func (c *Compute) stop() {
	c.cpu.stop()
	c.memory.stop()
	c.disk.stop()
}
