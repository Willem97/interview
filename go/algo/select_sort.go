package main

import (
	"fmt"
)

func main() {
	arr := []int{1, 3, 6, 8, 21, 24, 10}
	sortArr := selectSort(arr)
	fmt.Println(sortArr)
}

/*
	复杂度 O(n2)
	1.首先在未排序序列中找到最小（大）元素，存放到排序序列的起始位置
	2.再从剩余未排序元素中继续寻找最小（大）元素，然后放到已排序序列的末尾。
	3.重复第二步，直到所有元素均排序完毕。
*/
func selectSort(arr []int) []int {
	if len(arr) == 0 {
		return arr
	}

	for i := 0; i < len(arr); i++ {
		min := i
		for j := i + 1; j < len(arr); j++ {
			if arr[j] < arr[min] {
				min = j
			}
			arr[i], arr[min] = arr[min], arr[i]
		}
	}
	return arr
}
