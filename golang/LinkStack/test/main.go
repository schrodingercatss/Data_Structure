package main

import (
	"LinkStack"
	"fmt"
)

func main() {
	lstk := LinkStack.NewLinkStack()
	lstk.Push(1)
	lstk.Push(2)
	lstk.Push(3)
	lstk.Push(4)

	for !lstk.IsEmpty() {
		fmt.Println(lstk.Pop())
	}
}
