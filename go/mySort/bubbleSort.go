package main

import "fmt"

func main() {
	var list = make([]int, 0)
	list = append(list, 6, 7, 5, 4, 3, 2, 9, 0)
	fmt.Println("排序前", list)
	res := bubbleV3(list)
	fmt.Println("排序后", res)

}

func bubble(list []int) []int {
	n := len(list)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-1-i; j++ {
			if list[j] > list[j+1] {
				tmp := list[j+1]
				list[j+1] = list[j]
				list[j] = tmp
			}
		}
	}
	return list
}

func bubbleV2(list []int) []int {
	n := len(list)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-1-i; j++ {
			if list[j] > list[j+1] {
				tmp := list[j]
				list[j] = list[j+1]
				list[j+1] = tmp
			}
		}
	}

	return list
}

func bubbleV3(list []int) []int {
	n := len(list)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-1-i; j++ {
			if list[j] > list[j+1] {
				tmp := list[j]
				list[j] = list[j+1]
				list[j+1] = tmp
			}
		}
	}
	return list
}
