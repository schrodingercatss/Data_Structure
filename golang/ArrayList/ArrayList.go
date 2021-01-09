package ArrayList

import (
	"errors"
	"fmt"
)

type List interface {
	Size() int                                      // 获取数组的大小
	Get(index int) (interface{}, error)             // 获取第i个位置上的元素
	Set(index int, newElement interface{}) error    // 修改数据
	Insert(index int, newElement interface{}) error // 插入数据
	Append(newElement interface{})                  // 追加元素
	Clear()                                         // 清空数组
	Delete(index int) error                         // 删除元素
	String() string                                 // 返回数组字符串
	Iterator() Iterator                             // 迭代器
}

// DataStruct: ArrayList
type ArrayList struct {
	dataStore []interface{}
	length    int // 数组长度
}

func NewArrayList() *ArrayList {
	list := new(ArrayList)                      // 初始结构体
	list.dataStore = make([]interface{}, 0, 10) // 开辟空间为10的数组
	list.length = 0
	return list
}

func (list *ArrayList) Size() int {
	return list.length
}

func (list *ArrayList) Get(index int) (interface{}, error) {
	if index >= list.length || index < 0 {
		return nil, errors.New("索引越界")
	}
	return list.dataStore[index], nil
}

func (list *ArrayList) Set(index int, newElement interface{}) error {
	if index >= list.length || index < 0 {
		return errors.New("索引越界")
	}
	list.dataStore[index-1] = newElement
	return nil
}

func (list *ArrayList) Insert(index int, newElement interface{}) error {
	if index > list.length || index < 0 {
		return errors.New("索引越界")
	}
	// 内存不足时进行分配
	if list.length == cap(list.dataStore) {
		// 如果第一个参数为0，则表明没有开辟内存，所以要定义成2 * list.length
		newDataStore := make([]interface{}, 2*list.length, 2*list.length)
		copy(newDataStore, list.dataStore)
		list.dataStore = newDataStore
	}
	list.dataStore = list.dataStore[:list.length+1] // 插入数据，内存移动一位
	for i := list.length; i > index; i-- {
		list.dataStore[i] = list.dataStore[i-1] // 从后往前移动
	}
	list.dataStore[index] = newElement
	list.length++
	return nil
}

func (list *ArrayList) Append(newElement interface{}) {
	list.dataStore = append(list.dataStore, newElement)
	list.length++
}

func (list *ArrayList) Clear() {
	list.dataStore = make([]interface{}, 0, 10)
	list.length = 0
}

func (list *ArrayList) Delete(index int) error {
	if index >= list.length || index < 0 {
		return errors.New("索引越界")
	}
	list.dataStore = append(list.dataStore[:index], list.dataStore[index+1:]...)
	list.length--
	return nil
}

func (list *ArrayList) String() string {
	return fmt.Sprint(list.dataStore)
}
