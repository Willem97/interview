package main

import (
	"fmt"
)

func main() {
	arr := []int{1, 3, 6, 8, 21, 24, 10}
	sortArr := bubbleSort(arr)
	fmt.Println(sortArr)
}

func bubbleSort(arr []int) []int {
	if len(arr) == 0 {
		return arr
	}

	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr); j++ {
			if arr[i] < arr[j] {
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
	}
	return arr
}
