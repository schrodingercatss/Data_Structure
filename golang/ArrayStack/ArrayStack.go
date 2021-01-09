package ArrayStack

import (
	"errors"
	"log"
)

type ArrayStack interface {
	Size() int                      // 获取栈长度
	Top() interface{}               // 获取栈顶元素
	Pop() interface{}               // 弹出栈顶元素
	Push(element interface{}) error // 向栈顶插入元素
	IsFull() bool                   // 判断是否栈满
	IsEmpty() bool                  // 判断是否栈空
	Clear()                         // 清空栈
}

type Stack struct {
	dataStore []interface{}
	capacity  int // 栈的最大长度
	length    int // 栈当前长度
}

func NewStack() *Stack {
	mystack := new(Stack)
	mystack.dataStore = make([]interface{}, 0, 16)
	mystack.length = 0
	mystack.capacity = 16
	return mystack
}

func (stack *Stack) Size() int {
	return stack.length
}

func (stack *Stack) Top() interface{} {
	if stack.IsEmpty() {
		log.Fatal("栈空")
	}
	return stack.dataStore[stack.length-1]
}

func (stack *Stack) Pop() interface{} {
	if stack.IsEmpty() {
		log.Fatal("栈空")
	}
	last := stack.dataStore[stack.length-1]
	stack.dataStore = stack.dataStore[:stack.length-1]
	stack.length--
	return last
}

func (stack *Stack) Push(element interface{}) error {
	if stack.IsFull() {
		return errors.New("栈满")
	}
	stack.dataStore = append(stack.dataStore, element)
	stack.length++
	return nil
}

func (stack *Stack) IsFull() bool {
	if stack.length >= stack.capacity {
		return true
	}
	return false
}

func (stack *Stack) IsEmpty() bool {
	return stack.length == 0
}

func (stack *Stack) Clear() {
	stack.dataStore = make([]interface{}, 0, 16)
	stack.length = 0
}
