package main

import "fmt"

func removeElement(nums []int ,val int) int  {
	n := len(nums)
	var left int = 0
	for right := 0;right<n;right++{
		if nums[right] != val{
			nums[left] = nums[right]
			left++
		}
	}
	return left
}

func main() {
	var nums = make([]int,0)
	var val int = 3

	nums = append(nums,3,3,2,4)

	res := removeElement(nums,val)
	fmt.Println(res)
}