package Sorting

func RadixSort(arr []int) []int {
	max := GetMax(arr)
	for i := 1; max/i > 0; i *= 10 {
		arr = CountingSort(arr)
	}
	return arr
}
