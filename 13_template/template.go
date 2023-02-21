package _3_template

import "fmt"

/*
模板方法模式：定义一个操作中的算法的骨架，而将一些步骤延迟到子类中，使得子类可以不改变一个算法的结构即可重定义该算法的某些特定步骤。
模板方法（Template Method）是一个比较简单的模式。它的主要思想是，定义一个操作的一系列步骤，对于某些暂时确定不下来的步骤，就留给子类去实现好了，这样不同的子类就可以定义出不同的步骤。
*/

type Cooker interface {
	Open()
	Fire()
	Cooke()
	OutFire()
	Close()
}

type CookMenu struct {
}

func (CookMenu) Open() {
	fmt.Println("打开开关")
}

func (CookMenu) Fire() {
	fmt.Println("开火")
}

// Cooke 做菜，交给具体的子类实现
func (CookMenu) Cooke() {
}

func (CookMenu) OutFire() {
	fmt.Println("关火")
}

func (CookMenu) Close() {
	fmt.Println("关闭开关")
}

//DoCook 封装具体步骤
func DoCook(cook Cooker) {
	cook.Open()
	cook.Fire()
	cook.Cooke()
	cook.OutFire()
	cook.Close()
}

// TomatoScrambledEggs 西红柿炒蛋
type TomatoScrambledEggs struct {
	CookMenu
}

func (*TomatoScrambledEggs) Cooke() {
	fmt.Println("做一个西红柿炒蛋")
}

// MapoTofu 麻婆豆腐
type MapoTofu struct {
	CookMenu
}

func (*MapoTofu) Cooke() {
	fmt.Println("做一个麻婆豆腐")
}
