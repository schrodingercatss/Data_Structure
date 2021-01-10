package Sorting

func BucketSort(arr []int) []int {
	n := len(arr)
	// 获取最大值
	max := GetMax(arr)
	tmp := make([]int, max+1)

	for _, value := range arr {
		tmp[value]++
	}
	res := make([]int, 0, n)
	for i, v := range tmp {
		for ; v > 0; v-- {
			res = append(res, i)
		}
	}
	return res
}
