package _2_factory

import "fmt"

//type IKeyBoard interface {
//	write(tt string)
//}

type HPKeyBoard struct {
}

func (h HPKeyBoard) write(tt string) {
	fmt.Println("HP 键盘接受写入指令：", tt)
}

type IMouse interface {
	click(tt string)
}

type HPMouse struct {
}

func (h HPMouse) click(tt string) {
	fmt.Println("HP 鼠标接受点击指令：", tt)
}

type IAccessoriesFactory interface {
	CreateKeyBoardFactory() IKeyBoard
	CreateMouseFactory() IMouse
}

type HPAccessoriesFactory struct {
}

func (h HPAccessoriesFactory) CreateKeyBoardFactory() IKeyBoard {
	return HPKeyBoard{}
}

func (h HPAccessoriesFactory) CreateMouseFactory() IMouse {
	return HPMouse{}
}
