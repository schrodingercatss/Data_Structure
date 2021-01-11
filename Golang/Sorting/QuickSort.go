package Sorting

func QuickSort(arr []int, l, r int) {
	if l >= r {
		return
	}
	x := arr[l]
	i, j := l, r
	for i < j {
		for arr[i] < x {
			i++
		}
		for arr[j] > x {
			j--
		}
		if i < j {
			arr[i], arr[j] = arr[j], arr[i]
		}
	}
	QuickSort(arr, l, j)
	QuickSort(arr, j+1, r)
}
