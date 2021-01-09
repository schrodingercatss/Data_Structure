package main

import (
	"ArrayStack"
	"errors"
	"fmt"
	"io/ioutil"
)

// 使用栈模拟斐波那契的递归过程
func test1() {
	mystack := ArrayStack.NewStack()
	mystack.Push(7)
	ans := 0
	for !mystack.IsEmpty() {
		data := mystack.Pop()
		if data == 1 || data == 2 {
			ans += 1
		} else {
			mystack.Push(data.(int) - 1)
			mystack.Push(data.(int) - 2)
		}
	}
	fmt.Println(ans)
}

// 递归遍历目录文件
func GetAllFiles(path string, files []string) ([]string, error) {
	read, err := ioutil.ReadDir(path) // 读取文件夹
	if err != nil {
		return files, errors.New("文件夹不可读取")
	}
	for _, file := range read {
		if file.IsDir() {
			fullDir := path + "\\" + file.Name() // 构造新的路径
			files = append(files, fullDir)
			files, _ = GetAllFiles(fullDir, files) // 文件夹递归处理
		} else {
			fullDir := path + "\\" + file.Name()
			files = append(files, fullDir)
		}
	}
	return files, nil
}

func TreeGetAllFiles(path string, files []string, level int) ([]string, error) {
	levelStr := ""
	for i := 1; i <= level; i++ {
		levelStr += "--"
	}
	read, err := ioutil.ReadDir(path) // 读取文件夹
	if err != nil {
		return files, errors.New("文件夹不可读取")
	}
	for _, file := range read {
		if file.IsDir() {
			fullDir := path + "\\" + file.Name() // 构造新的路径
			files = append(files, levelStr+fullDir)
			files, _ = TreeGetAllFiles(fullDir, files, level+1) // 文件夹递归处理
		} else {
			fullDir := path + "\\" + file.Name()
			files = append(files, levelStr+fullDir)
		}
	}
	return files, nil
}

func test2() {
	path := "E:\\6.824"
	files := []string{}
	files, _ = TreeGetAllFiles(path, files, 0)
	for i := 0; i < len(files); i++ {
		fmt.Println(files[i])
	}
}

func StackGetAllFile() {
	path := "E:\\6.824"
	files := []string{}
	myStack := ArrayStack.NewStack()
	myStack.Push(path)
	for !myStack.IsEmpty() {
		path = myStack.Pop().(string)
		files = append(files, path)
		read, _ := ioutil.ReadDir(path)
		for _, file := range read {
			if file.IsDir() {
				fullDir := path + "\\" + file.Name()
				files = append(files, fullDir)
				myStack.Push(fullDir)
			} else {
				fullDir := path + "\\" + file.Name()
				files = append(files, fullDir)
			}
		}
	}
	for i := 0; i < len(files); i++ {
		fmt.Println(files[i])
	}
}

func main() {
	test2()
}
