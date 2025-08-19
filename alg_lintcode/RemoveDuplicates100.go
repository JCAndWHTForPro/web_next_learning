package main

func removeDuplicates(nums []int) int {
	for i := 1; i < len(nums); {
		if nums[i] == nums[i-1] {
			nums = append(nums[:i], nums[i+1:]...)
		} else {
			i++
		}
	}
	return len(nums)
}
