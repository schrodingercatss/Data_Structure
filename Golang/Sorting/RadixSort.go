package Sorting

// 对32位整数进行排序，基数取r = 2^16 ，这样我们可以将32位整数当成r进制，低16位算低位，高16位算高位
func RadixSort(arr []int) {
	n := len(arr)
	tmp := make([]int, n)
	r := 65535 // 基数
	L, H := make([]int, r+1), make([]int, r+1)
	// 计数
	for i := 0; i < n; i++ {
		L[arr[i]&r]++
		H[(arr[i]>>16)&r]++
	}
	// 求前缀和，求完前缀和后，这时每个（原来存在的）数i对应的L[i]，就是排序后最后一个i的位置加上1。
	//比如：num & r = a, 则L[a] - 1或者H[a] - 1就是该数因为存放的位置
	for i := 1; i <= r; i++ {
		L[i] += L[i-1]
		H[i] += H[i-1]
	}
	// 对低位进行计数排序
	for i := n - 1; i >= 0; i-- {
		L[arr[i]&r]--
		tmp[L[arr[i]&r]] = arr[i]
	}
	// 对高位进行计数排序
	for i := n - 1; i >= 0; i-- {
		H[(tmp[i]>>16)&r]--
		arr[H[(tmp[i]>>16)&r]] = tmp[i]
	}
}
