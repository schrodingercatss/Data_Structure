package main

import (
	"CircularQueue"
	"fmt"
)

func main() {
	cQueue := CircularQueue.NewCircularQueue()
	cQueue.Push(1)
	cQueue.Push(2)
	cQueue.Push(3)
	cQueue.Push(4)
	cQueue.Push(5)

	for i := 0; i < 5; i++ {
		fmt.Println(cQueue.Front())
		cQueue.Pop()
	}
	fmt.Println(cQueue.IsEmpty(), cQueue.IsFull())
}
