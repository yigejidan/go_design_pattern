package _2_factory

/*
工厂模式
*/

import "fmt"

type IKeyBoard interface {
	write(tt string)
}

type HPNormalKeyBoard struct {
}

func (h HPNormalKeyBoard) write(tt string) {
	fmt.Println("HP 普通键盘接受写入指令：", tt)
}

type LenovoNormalKeyBoard struct {
}

func (l LenovoNormalKeyBoard) write(tt string) {
	fmt.Println("Lenovo 普通键盘接受写入指令：", tt)
}

type HPMechanicalKeyBoard struct {
}

func (h HPMechanicalKeyBoard) write(tt string) {
	fmt.Println("HP 机械键盘接受写入指令：", tt)
}

type LenovoMechanicalKeyBoard struct {
}

func (h LenovoMechanicalKeyBoard) write(tt string) {
	fmt.Println("Lenovo 机械键盘接受写入指令：", tt)
}

type IKeyBoardFactory interface {
	Create(keyBoardType string) IKeyBoard
}

type HPKeyBoardFactory struct {
}

func (h HPKeyBoardFactory) Create(keyBoardType string) IKeyBoard {
	switch keyBoardType {
	case "Normal":
		return HPNormalKeyBoard{}
	case "Mechanical":
		return HPMechanicalKeyBoard{}
	default:
		panic("未知键盘类型")
	}
}

type LenovoKeyBoardFactory struct {
}

func (h LenovoKeyBoardFactory) Create(keyBoardType string) IKeyBoard {
	switch keyBoardType {
	case "Normal":
		return LenovoNormalKeyBoard{}
	case "Mechanical":
		return LenovoMechanicalKeyBoard{}
	default:
		panic("未知键盘类型")
	}
}

func NewKeyBoardFactory(brand string) IKeyBoardFactory {
	switch brand {
	case "HP":
		return HPKeyBoardFactory{}
	case "Lenovo":
		return LenovoKeyBoardFactory{}
	default:
		panic("未知工厂厂商")
	}
}
