// 代理模式
/*
代理模式是一种软件设计模式，属于结构型模式。它通过创建一个代理对象，可以控制对其他对象的访问，并在访问对象时加入一些额外的操作。

代理模式的主要目标是通过引入一个代理对象，将客户端和目标对象解耦，使得代理对象可以在不改变目标对象的情况下为其提供额外的功能或控制访问。

在代理模式中，通常有以下角色：

抽象主题（Subject）: 定义了目标对象和代理对象共同实现的接口。
真实主题（Real Subject）: 实现了抽象主题接口，是业务逻辑的具体实现对象。
代理（Proxy）: 实现了抽象主题接口，并持有一个对真实主题对象的引用。它在访问真实主题之前或之后进行一些额外的操作。
代理模式有多种类型，包括静态代理、动态代理和虚拟代理等。

静态代理是在编译时就已经确定代理类和真实类的关系，代理类和真实类之间是一对一的关系。编写代码时需要为每个代理类手动编写代码，适用于一些简单的场景。
动态代理是在运行时通过反射机制动态生成代理类，代理类和真实类之间是一对多的关系。可以通过动态代理来代理一组接口相似的对象，减少了重复编写代理类的工作，适用于一些复杂的场景。
虚拟代理延迟创建真实对象，只有在需要的时候才创建，可以减少系统的开销。
代理模式的主要优点包括：

隔离性：代理对象可以将客户端与目标对象隔离开，保护目标对象的安全性。
扩展性：代理对象可以通过继承或实现抽象主题接口来灵活地进行功能扩展。
权限控制：代理对象可以在访问目标对象时添加权限控制的逻辑。
懒加载：虚拟代理可以延迟创建目标对象，提高系统的性能。
总结来说，代理模式提供了一种控制对其他对象的访问的方法，通过引入代理对象来为目标对象提供额外的功能或控制访问。它是一种很有用的模式，常用于远程代理、虚拟代理、权限控制等场景。
*/
package main

import (
	"errors"
	"fmt"
	"io/ioutil"
)

type NewSubject interface {
	getContent() (string, error)
	setContent(content string) error
}

type RealSubject struct {
	filename string
	content  string
}

func (r *RealSubject) getContent() (string, error) {
	if r.content == "" {
		data, err := ioutil.ReadFile(r.filename)
		if err != nil {
			return "", err
		}
		r.content = string(data)
	}
	return r.content, nil
}

func (r *RealSubject) setContent(content string) error {
	err := ioutil.WriteFile(r.filename, []byte(content), 0666)
	if err != nil {
		return err
	}
	r.content = content
	return nil
}

func NewRealSubject(filename string) *RealSubject {
	return &RealSubject{filename: filename}
}

type VirtualProxy struct {
	filename string
	sub      NewSubject
}

func NewVirtualProxy(filename string) *VirtualProxy {
	return &VirtualProxy{filename: filename}
}

func (v *VirtualProxy) getContent() (string, error) {
	if v.sub == nil {
		v.sub = NewRealSubject(v.filename)
	}
	return v.sub.getContent()
}

func (v *VirtualProxy) setContent(content string) error {
	if v.sub == nil {
		v.sub = NewRealSubject(v.filename)
	}
	return v.sub.setContent(content)
}

type ProtectProxy struct {
	sub NewSubject
}

func NewProtectProxy(filename string) *ProtectProxy {
	return &ProtectProxy{sub: NewRealSubject(filename)}
}

func (p *ProtectProxy) getContent() (string, error) {
	return p.sub.getContent()
}

func (p *ProtectProxy) setContent(content string) error {
	return errors.New("没有权限")
}

func main() {
	file_name := "/Users/yjz/code/gocode/golearnpro/design_pattern/test.txt"
	sub := NewRealSubject(file_name)
	content, err := sub.getContent()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf(content)
	}

	virtProxy := NewVirtualProxy(file_name)
	content, err = virtProxy.getContent()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(content)
	}

	protectProxy := NewProtectProxy(file_name)
	content, err = protectProxy.getContent()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(content)
		err = protectProxy.setContent("11")
		if err != nil {
			fmt.Println(err)
		}
	}

}
