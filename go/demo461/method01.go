package main

import (
	"fmt"
	"math/bits"
)

func main() {
	var x int  = 3
	var y int  = 2
	var z int
	z = hammingDistance(x,y)
	fmt.Println(z)
}

//461. 汉明距离
func hammingDistance(x int, y int) int {
	return bits.OnesCount(uint(x ^ y))
}