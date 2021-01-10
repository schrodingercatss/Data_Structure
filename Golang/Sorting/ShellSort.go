package Sorting

func ShellSort(arr []int) {
	n := len(arr)
	for step := n / 2; step > 0; step /= 2 {
		for i := step; i < n; i++ {
			tmp := arr[i]
			var j int
			for j = i; j >= step; j -= step {
				if tmp < arr[j-step] {
					arr[j] = arr[j-step]
				} else {
					break
				}
			}
			arr[j] = tmp
		}
	}
}
