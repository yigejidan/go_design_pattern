package _4_prototype

/*
原型模式
*/

import (
	"encoding/json"
	"time"
)

type Keyword struct {
	word      string
	visit     int
	UpdatedAt *time.Time
}

func (k *Keyword) Clone() *Keyword {
	var newKeyWord Keyword
	b, _ := json.Marshal(k)
	err := json.Unmarshal(b, &newKeyWord)
	if err != nil {
		return nil
	}
	return &newKeyWord
}

type KeyWords map[string]*Keyword

func (words KeyWords) Clone(updatedWords []*Keyword) KeyWords {
	newKeyWords := KeyWords{}

	for k, v := range words {
		// 这里是浅拷贝，直接拷贝了地址
		newKeyWords[k] = v
	}

	// 替换掉需要更新的字段，这里用的是深拷贝
	for _, word := range updatedWords {
		newKeyWords[word.word] = word.Clone()
	}
	return newKeyWords
}
