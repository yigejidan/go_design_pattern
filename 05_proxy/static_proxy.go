package _5_proxy

/*
静态代理模式
*/

import (
	"log"
	"time"
)

type IUser interface {
	Login(username, password string) error
}

type User struct {
}

func (u *User) Login(username, password string) error {
	return nil
}

type userProxy struct {
	user *User
}

func NewUserProxy(user *User) *userProxy {
	return &userProxy{
		user: user,
	}
}

func (p *userProxy) Login(username, password string) error {
	start := time.Now()

	if err := p.user.Login(username, password); err != nil {
		return err
	}
	// after 这里可能也有一些监控统计的逻辑
	log.Printf("user login cost time: %s", time.Now().Sub(start))

	return nil
}
