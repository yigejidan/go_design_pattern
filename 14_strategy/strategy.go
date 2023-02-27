package _4_strategy

import (
	"fmt"
	"io/ioutil"
	"os"
)

/*
策略模式：定义一系列的算法，把它们一个个封装起来，并且使它们可相互替换。本模式使得算法可独立于使用它的客户而变化。
策略模式：Strategy，是指，定义一组算法，并把其封装到一个对象中。然后在运行时，可以灵活的使用其中的一个算法。
*/

// StorageStrategy 存储策略
type StorageStrategy interface {
	Save(name string, data []byte) error
}

// fileStorage 保存到文件
type fileStorage struct{}

func (s *fileStorage) Save(name string, data []byte) error {
	return ioutil.WriteFile(name, data, os.ModeAppend)
}

// encryptFileStorage 保存到加密文件
type encryptFileStorage struct{}

func (e *encryptFileStorage) Save(name string, data []byte) error {
	data, err := encrypt(data)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(name, data, os.ModeAppend)
}

func encrypt(data []byte) ([]byte, error) {
	// 加密算法
	return data, nil
}

// 将一组算法封装到map中
var strategys = map[string]StorageStrategy{
	"file":         &fileStorage{},
	"encrypt_file": &encryptFileStorage{},
}

// 使用时根据灵活选择算法
func NewStorageStrategy(t string) (StorageStrategy, error) {
	s, ok := strategys[t]
	if !ok {
		return nil, fmt.Errorf("not found StorageStrategy: %s", t)
	}
	return s, nil
}
