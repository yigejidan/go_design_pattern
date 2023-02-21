package _8_adapter

import "fmt"

type ICharge interface {
	ChongDian() error
}

type TypeC struct {
}

func (t *TypeC) chaTou() error {
	fmt.Println("TypeC 接口充电")
	return nil
}

type Anzhou struct {
}

func (a *Anzhou) chaTou() error {
	fmt.Println("安卓接口充电")
	return nil
}

type TypeCCharge struct {
	client TypeC
}

func (t *TypeCCharge) ChongDian() error {
	return t.client.chaTou()
}

type AnzhouCharge struct {
	client Anzhou
}

func (a *AnzhouCharge) ChongDian() error {
	return a.client.chaTou()
}
