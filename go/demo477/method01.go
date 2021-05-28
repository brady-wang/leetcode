package main

import "fmt"

func main() {
	var nums = make([]int,0)
	nums = append(nums,4,3)
	// 11101
	fmt.Println(29 >> 2 )

	total := totalHammingDistance(nums)
	fmt.Println(total)
}

//数组中元素的范围为从 0到 10^9。  由于 10^9 < 2^30 我们可以直接从二进制的第 000 位枚举到第 292929 位。
//数组的长度不超过 10^

// 假设把每个数组里面的数字都是30位的,找到他们每个数对应 第一位 第二位 第三位等的1的数量,0的数量  相乘就是不同的,汉明距离  所以外层循环为30 因为
func totalHammingDistance(nums []int) (total int) {
	n := len(nums)
	for i:=0;i<30;i++{
		c := 0
		for _,value := range nums{
			c += value >> i & 1

		}

		total += c*(n-c)
	}
	return 
}