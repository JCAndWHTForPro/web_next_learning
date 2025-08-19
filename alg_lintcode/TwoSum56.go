package main

/*
*描述
给一个整数数组，找到两个数使得他们的和等于一个给定的数 target。

你需要实现的函数twoSum需要返回这两个数的下标, 并且第一个下标小于第二个下标。注意这里下标的范围是 0 到 n-1。

你可以假设只有一组答案。

数组内元素不能重复使用
*/
func TwoSum(numbers []int, target int) []int {
	// write your code here
	i := len(numbers)
	if i == 0 {
		return []int{}
	}
	var mapping = make(map[int]int)
	for index, value := range numbers {
		mapping[target-value] = index
	}
	for index, value := range numbers {
		if v, ok := mapping[value]; ok && index != v {
			if index > v {
				return []int{v, index}
			}
			return []int{index, v}
		}

	}
	return []int{}
}
