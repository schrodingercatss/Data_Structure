package Sorting

func HeapSort(arr []int) {
	n := len(arr)
	for i := n/2 - 1; i >= 0; i-- {
		HeapAdjust(arr, n, i)
	}
	for i := n - 1; i > 0; i-- {
		arr[0], arr[i] = arr[i], arr[0]
		HeapAdjust(arr, i, 0)
	}
}

func HeapAdjust(arr []int, n int, index int) {
	leftChild := 2*index + 1
	rightChild := 2*index + 2
	max := index
	if leftChild < n && arr[leftChild] > arr[max] {
		max = leftChild
	}
	if rightChild < n && arr[rightChild] > arr[max] {
		max = rightChild
	}
	if index != max {
		arr[index], arr[max] = arr[max], arr[index]
		HeapAdjust(arr, n, max)
	}
}
