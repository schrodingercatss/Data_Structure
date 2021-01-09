package ArrayQueue

import "log"

type Queue interface {
	Front() interface{}       // 返回队首元素
	Push(element interface{}) // 入队
	Pop() interface{}         // 出队
	Size() int                // 返回队列长度
	IsEmpty() bool            // 判空
	IsFull() bool             // 判满
}

type ArrayQueue struct {
	length    int
	capacity  int
	dataStore []interface{}
}

func NewArrayQueue() *ArrayQueue {
	queue := new(ArrayQueue)
	queue.length, queue.capacity = 0, 100
	queue.dataStore = make([]interface{}, 0, 100)
	return queue
}

func (queue *ArrayQueue) Size() int {
	return queue.length
}

func (queue *ArrayQueue) IsEmpty() bool {
	return queue.length == 0
}

func (queue *ArrayQueue) IsFull() bool {
	return queue.length == queue.capacity
}

func (queue *ArrayQueue) Front() interface{} {
	if queue.IsEmpty() {
		return nil
	}
	return queue.dataStore[0]
}

func (queue *ArrayQueue) Push(element interface{}) {
	if queue.IsFull() {
		log.Fatal("队满")
	}
	queue.dataStore = append(queue.dataStore, element)
	queue.length++
}

func (queue *ArrayQueue) Pop() interface{} {
	if queue.IsEmpty() {
		return nil
	}
	res := queue.dataStore[0]
	queue.length--
	queue.dataStore = append(queue.dataStore[1:])
	return res
}
