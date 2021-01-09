package main

import (
	"ArrayQueue"
	"fmt"
	"io/ioutil"
)

func test() {
	queue := ArrayQueue.NewArrayQueue()
	queue.Push(1)
	queue.Push(2)
	queue.Push(3)
	queue.Push(4)

	for i := 0; i < 4; i++ {
		fmt.Println(queue.Front())
		queue.Pop()
	}

}

// 队列遍历目录下所有文件
func main() {
	path := "E:\\6.824"
	files := []string{}

	queue := ArrayQueue.NewArrayQueue()
	queue.Push(path)

	for !queue.IsEmpty() {
		path := queue.Pop().(string)
		read, _ := ioutil.ReadDir(path)
		for _, file := range read {
			fullDir := path + "\\" + file.Name()
			if file.IsDir() {
				fmt.Println("Dir", fullDir)
				queue.Push(fullDir)
			} else {
				files = append(files, fullDir)
				fmt.Println("File", fullDir)
			}
		}
	}
}
