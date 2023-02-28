package _0_command

import "fmt"

/*
命令模式：将一个请求封装为一个对象，从而使你可用不同的请求对客户进行参数化，对请求排队或记录请求日志，以及支持可撤销的操作。
命令模式（Command）是指，把请求封装成一个命令，然后执行该命令。
*/

// ICommand 命令接口
type ICommand interface {
	Execute() error
}

// StartCommand 游戏开始运行
type StartCommand struct {
}

// NewStartCommand new一个开始游戏命令
func NewStartCommand() *StartCommand {
	return &StartCommand{}
}

func (s *StartCommand) Execute() error {
	fmt.Println("游戏开始运行")
	return nil
}

type ArchiveCommand struct {
}

func NewArchiveCommand() *ArchiveCommand {
	return &ArchiveCommand{}
}

func (a *ArchiveCommand) Execute() error {
	fmt.Println("游戏开始存档")
	return nil
}
