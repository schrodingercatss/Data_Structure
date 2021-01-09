package main

import (
	"LinkQueue"
	"fmt"
)

func main() {
	lqueue := LinkQueue.NewLinkQueue()
	lqueue.Push(1)
	lqueue.Push(2)
	lqueue.Push(3)
	lqueue.Push(4)
	fmt.Println("Size: ", lqueue.Size())

	for !lqueue.IsEmpty() {
		fmt.Println(lqueue.Pop())
	}

	fmt.Println(lqueue.IsEmpty())
}
