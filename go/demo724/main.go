package main

import "fmt"

func main() {
	var list = make([]int, 0)

	list = append(list, 1, 7, 3, 6, 5, 6)

	index := pivotIndex(list)
	fmt.Println(index)

}

func pivotIndex(nums []int) int {
	n := len(nums)

	var total int = 0
	var sum int = 0
	for _, v := range nums {
		total = total + v
	}

	for i := 0; i < n; i++ {
		if 2*sum+nums[i] == total {
			return i
		}
		sum = sum + nums[i]
	}

	return -1
}
