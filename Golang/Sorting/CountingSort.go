package Sorting

// arr切片存储待排数据，k为数据最大值，排序范围为0~k
func CountingSort(arr []int) []int {
	n := len(arr)
	// 获取n个数的最大值
	max, min := GetMax(arr), GetMin(arr)
	count := make([]int, max-min+1)
	res := make([]int, n)
	for _, num := range arr {
		count[num-min]++
	}
	index := 0
	for i := 0; i < len(count); i++ {
		for ; count[i] > 0; count[i]-- {
			res[index] = i + min
			index++
		}
	}
	return res
}
