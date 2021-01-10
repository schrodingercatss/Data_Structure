package CircularQueue

import "log"

type Queue interface {
	Front() interface{}       // 返回队首元素
	Push(element interface{}) // 入队
	Pop() interface{}         // 出队
	Size() int                // 返回队列长度
	IsEmpty() bool            // 判空
	IsFull() bool             // 判满
}

const QUEUE_SIZE = 100

type CircularQueue struct {
	dataStore   [QUEUE_SIZE]interface{}
	front, rear int
}

func NewCircularQueue() *CircularQueue {
	cQueue := new(CircularQueue)
	cQueue.front, cQueue.rear = 0, 0
	return cQueue
}

func (q *CircularQueue) Size() int {
	return (q.rear - q.front + QUEUE_SIZE) % QUEUE_SIZE
}

func (q *CircularQueue) IsEmpty() bool {
	return q.front == q.rear
}

func (q *CircularQueue) IsFull() bool {
	return ((q.rear + 1) % QUEUE_SIZE) == q.front
}

func (q *CircularQueue) Push(element interface{}) {
	if q.IsFull() {
		log.Fatal("队满")
	}
	q.dataStore[q.rear] = element
	q.rear = (q.rear + 1) % QUEUE_SIZE
}

func (q *CircularQueue) Front() interface{} {
	if q.IsEmpty() {
		return nil
	}
	return q.dataStore[q.front]
}

func (q *CircularQueue) Pop() interface{} {
	if q.IsEmpty() {
		return nil
	}
	res := q.dataStore[q.front]
	q.front = (q.front + 1) % QUEUE_SIZE
	return res
}
