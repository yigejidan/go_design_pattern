package _9_facade

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUserService_LoginOrRegister(t *testing.T) {
	service := UserService{
		loginService:    loginService{},
		registerService: registerService{},
	}
	user, err := service.LoginOrRegister(13001010101, 1234)
	assert.NoError(t, err)
	assert.Equal(t, &User{Name: "test login"}, user)
}
