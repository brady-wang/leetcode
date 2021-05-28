package main

import (
	"fmt"
	"math/bits"
)

func main() {
	var x int  = 3 // 11
	var y int  = 2 // 10
	var z int
	z = hammingDistance(x,y)
	fmt.Println(z)
}

// 461. 汉明距离
// 计算 xxx 和 yyy 之间的汉明距离，可以先计算 x⊕y，然后统计结果中等于 1 的位数。
func hammingDistance(x int, y int) int {
	return bits.OnesCount(uint(x ^ y))
}