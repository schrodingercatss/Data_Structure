package ArrayList

import "errors"

type Iterator interface {
	HasNext() bool              // 是否存在下一个
	Next() (interface{}, error) // 获取下一个
	Remove()                    // 删除
	GetIndex() int              // 返回当前索引
}

type Iterable interface {
	Iterator() Iterator // 构造初始化接口
}

type ArrayListIterator struct {
	list         *ArrayList // 数组指针
	currentIndex int        // 当前索引
}

func (it *ArrayListIterator) HasNext() bool {
	return it.currentIndex < it.list.length
}

func (it *ArrayListIterator) Next() (interface{}, error) {
	if !it.HasNext() {
		return nil, errors.New("没有下一个")
	}
	value, err := it.list.Get(it.currentIndex)
	it.currentIndex++
	return value, err
}

func (it *ArrayListIterator) Remove() {
	it.currentIndex--
	it.list.Delete(it.currentIndex)
}

func (it *ArrayListIterator) GetIndex() int {
	return it.currentIndex
}

// 返回一个迭代器
func (list *ArrayList) Iterator() Iterator {
	it := new(ArrayListIterator)
	it.currentIndex = 0
	it.list = list
	return it
}
