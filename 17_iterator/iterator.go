package _7_iterator

/*
迭代器模式：提供一种方法顺序访问一个聚合对象中的各个元素，而又不需要暴露该对象的内部表示。
迭代器模式（Iterator）实际上在Java的集合类中已经广泛使用了。
我们以List为例，要遍历ArrayList，即使我们知道它的内部存储了一个Object[]数组，也不应该直接使用数组索引去遍历，因为这样需要了解集合内部的存储结构。
如果使用Iterator遍历，那么，ArrayList和LinkedList都可以以一种统一的接口来遍历
*/

// Iterator 迭代器接口
type Iterator interface {
	HasNext() bool
	Next()
	//CurrentItem 获取当前元素
	CurrentItem() interface{}
}

// ArrayInt 数组
type ArrayInt []int

// Iterator 返回迭代器
func (a ArrayInt) Iterator() Iterator {
	return &ArrayIntIterator{
		arrayInt: a,
		index:    0,
	}
}

// ArrayIntIterator 数组迭代
type ArrayIntIterator struct {
	arrayInt ArrayInt
	index    int
}

func (a *ArrayIntIterator) HasNext() bool {
	return a.index < len(a.arrayInt)-1
}

func (a *ArrayIntIterator) Next() {
	a.index++
}

func (a *ArrayIntIterator) CurrentItem() interface{} {
	return a.arrayInt[a.index]
}
