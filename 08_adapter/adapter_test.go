package _8_adapter

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAliyunClientAdapter_CreateServer(t *testing.T) {
	// 确保 adapter 实现了目标接口
	var a ICreateServer = &ALiYunClientAdapter{
		Client: ALiYunClient{},
	}

	assert.Nil(t, a.CreateServer(1.0, 2.0))

}

func TestAwsClientAdapter_CreateServer(t *testing.T) {
	// 确保 adapter 实现了目标接口
	var a ICreateServer = &AwsClientAdapter{
		Client: AWSClient{},
	}

	assert.Nil(t, a.CreateServer(1.0, 2.0))
}
