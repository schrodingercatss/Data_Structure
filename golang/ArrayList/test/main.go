package main

import (
	"ArrayList"
	"fmt"
)

func main() {
	var list ArrayList.List = ArrayList.NewArrayList()
	list.Append("A1")
	list.Append("B2")
	list.Append("C3")
	list.Append("D4")
	list.Append("E5")

	for it := list.Iterator(); it.HasNext(); {
		item, _ := it.Next()
		if item == "C3" {
			it.Remove()
		}
		fmt.Println(item)
	}
	fmt.Println(list)

}
