package main

import "fmt"

func main() {
	var list = make([]int, 0)
	list = append(list, 23, 42, 12, 34, 123, 45, 5, 2, 12)
	fmt.Println("排序前", list)
	res := selectSort(list)
	fmt.Println("排序后", res)

}

func selectSort(list []int ) []int{
	n := len(list)
	for i:=0; i<n-1;i++{
		min := i
		for j:=i+1;j<n;j++{
			if list[j] < list[min]{
				tmp:= list[j]
				list[j] = list[min]
				list[min] = tmp
			}
		}
	}
	return list
}
