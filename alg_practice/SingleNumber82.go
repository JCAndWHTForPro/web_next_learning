package main

/*
*
给出 2 * n + 1个数字，除其中一个数字之外其他每个数字均出现两次，找到这个数字。
n≤100
挑战
一次遍历，常数级的额外空间复杂度
*/
func SingleNumber(a []int) int {
	// write your code here
	var rst = 0
	for _, v := range a {
		rst ^= v
	}
	return rst
}
