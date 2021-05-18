package main

import "fmt"

func main() {
	var list = make([]int, 0)
	list = append(list, 23, 42, 12, 34, 123, 45, 5, 2, 12)
	fmt.Println("排序前", list)
	res := bubble(list)
	fmt.Println("排序后", res)

}




// 外层循环控制要走多上趟 ，内层从第一个开始依次与后面的对比，如果大就交换
// 内存循环次数，为外层循环完后，就可以少循环几次，因为每次都会有最大的在最后面了，不需要比较了
func bubble(list []int) []int {
	length := len(list)
	for i := 0; i < length-1; i++ {
		fmt.Println(list)
		for j := 0; j < length-1-i; j++ {
			if list[j] > list[j+1] {
				tmp := list[j]
				list[j] = list[j+1]
				list[j+1] = tmp
			}
		}
	}
	return list
}
