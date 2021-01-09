package Sorting

func BubbleSort(arr []int) {
	n := len(arr)
	flag := true
	for i := 0; i < n && flag; i++ {
		flag = false
		for j := 0; j < n-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
				flag = true
			}
		}
	}
}
