package main

import "fmt"

func main() {
	var list = make([]int, 0)
	list = append(list, 4,5,3,2,7,1,9,0,8)
	fmt.Println("排序前", list)
	res := insertSortV2(list)
	fmt.Println("排序后", res)

}


//思路：通过构建有序序列，对于未排序数据，在已排序序列中从后向前扫描，找到相应位置并插入。可以理解为玩扑克牌时的理牌；
func insertSort(list []int) []int {
	n := len(list)
	for i:=1;i<n;i++ {
		tmp := list[i]

		// 从后到前面 依次拿tmp和他们比较，如果比他们小 原有往后移动，tmp占用他的位置
		for j:=i-1;j>=0;j--{
			fmt.Println(j)
			fmt.Println(list)
			if tmp < list[j] {
				list[j+1] = list[j]
				list[j] = tmp
			}else {
				//break
			}
		}
		fmt.Println("-----------")
	}
	return list
}

func insertSortV2(list []int) [] int  {
	var n int
	n = len(list)

	for i:=1;i<n;i++{
		tmp := list[i]

		for j:=i-1;j>=0;j--{
			if tmp < list[j] { // 当前位置往后移，空出来给要插入的让位置 因为j是从i-1开始的，所以最后肯定是多了一个位置
				list[j+1] = list[j]
				list[j] = tmp
			} else {
				break // 因为对于已排序区，最后一个就是最大的了，你比最大的还大了，没必要再比较更小的了
			}
		}
	}

	return list
}