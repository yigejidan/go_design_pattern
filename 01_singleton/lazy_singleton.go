package _1_singleton

import "sync"

var (
	lazySingleton *Singleton
	once          = &sync.Once{}
)

// GetLazySingleton 懒汉式单例模式
func GetLazySingleton() *Singleton {
	if lazySingleton == nil {
		once.Do(func() {
			lazySingleton = &Singleton{}
		})
	}
	return lazySingleton
}
