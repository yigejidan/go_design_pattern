package _1_singleton

// Singleton 饿汉式单例
type Singleton struct{}

var singleton *Singleton

func init() {
	singleton = &Singleton{}
}

func GetSingleton() *Singleton {
	return singleton
}
