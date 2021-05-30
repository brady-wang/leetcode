package main

import "fmt"

func main() {
	var nums = make([]int,0)
	nums = append(nums, 2,7,11,15)
	var target int = 9
	res := twoSum(nums,target)
	fmt.Println(res)
}

func twoSum(nums []int, target int) []int {

	var hashTable = make(map[int]int, 0)
	if len(nums) == 0 {
		return nil
	}

	for k, x := range nums {
		if p, ok := hashTable[target-x]; ok {
			return []int{p, k}
		}
		hashTable[x] = k
	}
	return nil

}
