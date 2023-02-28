package _2_mediator

import (
	"fmt"
	"reflect"
)

/*
中介模式：用一个中介对象来封装一系列的对象交互。中介者使各个对象不需要显式地相互引用，从而使其耦合松散，而且可以独立地改变它们之间的交互。
中介模式（Mediator）又称调停者模式，它的目的是把多方会谈变成双方会谈，从而实现多方的松耦合。
*/

type Input string

func (i Input) String() string {
	return string(i)
}

type Selection string

func (s *Selection) Selected() string {
	return string(*s)
}

type Button struct {
	onClick func()
}

func (b *Button) SetOnClick(f func()) {
	b.onClick = f
}

type IMediator interface {
	HandleEvent(component interface{})
}

type Dialog struct {
	LoginButton         *Button
	RegButton           *Button
	Selection           *Selection
	UsernameInput       *Input
	PasswordInput       *Input
	RepeatPasswordInput *Input
}

func (d *Dialog) HandleEvent(component interface{}) {
	switch {
	case reflect.DeepEqual(component, d.Selection):
		if d.Selection.Selected() == "登录" {
			fmt.Println("select login")
			fmt.Printf("show: %s\n", d.UsernameInput)
			fmt.Printf("show: %s\n", d.PasswordInput)
		} else if d.Selection.Selected() == "注册" {
			fmt.Println("select register")
			fmt.Printf("show: %s\n", d.UsernameInput)
			fmt.Printf("show: %s\n", d.PasswordInput)
			fmt.Printf("show: %s\n", d.RepeatPasswordInput)
		}
		// others, 如果点击了登录按钮，注册按钮
	}
}
