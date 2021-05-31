package main

import "fmt"

func main() {
	var list = make([]int, 0)
	list = append(list, 4,5,3,2,7,1,9,0,8)
	fmt.Println("排序前", list)
	res := quickSort(list)
	fmt.Println("排序后", res)

}

func quickSort(list []int  ) []int   {

	return list
}