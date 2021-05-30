package main

import "fmt"

func main() {

	var list = make([]string,0)
	list = append(list,"sdfsdf","sd","sdf")

	res := longestCommonPrefix(list)
	fmt.Println(res)

}

func longestCommonPrefix(strs []string) string {
	length := len(strs)
	if length == 0{
		return ""
	}
	for i:=0;i<len(strs[0]);i++{
		for j:=1; j<length;j++{
			fmt.Println(i,len(strs[j]))
			// 条件1 后面字符串已经到头了 比如 当前sdf 下标2 后面的sd为2 说明第一个的串已经超出了 取0到他前面的就是相同的
			// 条件2 发现不同的
			if i == len(strs[j]) || strs[j][i] != strs[0][i]{
				return strs[0][:i]
			}
		}
	}
	return strs[0]
}