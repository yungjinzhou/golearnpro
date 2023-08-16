// 抽象工厂
/*

 */

package main

import "fmt"

// 抽象产品

type PhoneShell interface {
	showShell()
}

type PhoneOS interface {
	showOS()
}

type PhoneCpu interface {
	showCpu()
}

// 抽象工厂

type PhoneFactory interface {
	makeShell() PhoneShell
	makeCpu() PhoneCpu
	makeOs() PhoneOS
}

// 具体产品

type MiShell struct {
}

type AppleShell struct {
}

type IOS struct {
}

type AndroidOS struct {
}

type SnapDragonCPU struct {
}

type AppleCPU struct {
}

func (m *MiShell) showShell() {
	fmt.Printf("小米手机壳")
}

func (a *AppleShell) showShell() {
	fmt.Printf("苹果手机壳")
}

func (a *AndroidOS) showOS() {
	fmt.Printf("安卓系统")
}

func (i *IOS) showOS() {
	fmt.Printf("苹果系统")
}

func (s *SnapDragonCPU) showCpu() {
	fmt.Printf("高通骁龙cpu")
}

func (a *AppleCPU) showCpu() {
	fmt.Printf("苹果cpu")
}

// 具体产品工厂

type MiFactory struct {
}

func (m *MiFactory) makeShell() PhoneShell {
	return &MiShell{}
}

func (m *MiFactory) makeCpu() PhoneCpu {
	return &SnapDragonCPU{}
}

func (m *MiFactory) makeOs() PhoneOS {
	return &AndroidOS{}
}

type AppleFactory struct {
}

func (m *AppleFactory) makeShell() PhoneShell {
	return &AppleShell{}
}

func (m *AppleFactory) makeCpu() PhoneCpu {
	return &AppleCPU{}
}

func (m *AppleFactory) makeOs() PhoneOS {
	return &IOS{}
}

// 客户端

type Phone struct {
	cpu   PhoneCpu
	shell PhoneShell
	os    PhoneOS
}

func (p *Phone) showInfo() {
	p.cpu.showCpu()
	p.os.showOS()
	p.shell.showShell()
}

func make_phone(phoneFactory PhoneFactory) *Phone {
	cpu := phoneFactory.makeCpu()
	os := phoneFactory.makeOs()
	shell := phoneFactory.makeShell()
	return &Phone{
		cpu:   cpu,
		os:    os,
		shell: shell,
	}
}

func main() {
	ph := make_phone(&MiFactory{})
	ph.showInfo()
}
