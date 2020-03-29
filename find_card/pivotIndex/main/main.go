package main

//https://leetcode-cn.com/problems/find-pivot-index/
func pivotIndex(nums []int) int {
	if len(nums) < 3{
		return -1
	}

	var lift int
	var right int
	var index = -1

	for i:= 1; i < len(nums); i++{
		right += nums[i]
	}

	for i := 0; i < len(nums); i++{
		if lift == right{
			index = i
			break
		}
		lift = lift + nums[i]
		if i+1 >= len(nums){
			right = 0
		}else {
			right = right - nums[i+1]
		}
	}
	return index

}