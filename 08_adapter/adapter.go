package _8_adapter

import "fmt"

/*
适配器模式：将一个类的接口转换成客户希望的另外一个接口，使得原本由于接口不兼容而不能一起工作的那些类可以一起工作。
*/

//ICreateServer 创建云主机接口
type ICreateServer interface {
	CreateServer(cpu, mem float64) error
}

//AWSClient AWS SDK
type AWSClient struct {
}

func (a *AWSClient) RunInstance(cpu, mem float64) error {
	fmt.Printf("aws client run success, cpu： %f, mem: %f", cpu, mem)
	return nil
}

// AwsClientAdapter AWS SDK 适配器
type AwsClientAdapter struct {
	Client AWSClient
}

func (a *AwsClientAdapter) CreateServer(cpu, mem float64) error {
	return a.Client.RunInstance(cpu, mem)
}

type ALiYunClient struct{}

func (a *ALiYunClient) RunInstance(cpu, mem float64) error {
	fmt.Printf("aliyun client run success, cpu： %f, mem: %f", cpu, mem)
	return nil
}

type ALiYunClientAdapter struct {
	Client ALiYunClient
}

func (a *ALiYunClientAdapter) CreateServer(cpu, mem float64) error {
	return a.Client.RunInstance(cpu, mem)
}
