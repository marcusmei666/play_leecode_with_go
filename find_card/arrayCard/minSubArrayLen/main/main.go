package main

func minSubArrayLen(s int, nums []int) int {
	var sum,k = 0,0

	for i := 0; i <= len(nums);  {
		sum += nums[i]
		if sum >= s{
			k++
			i++
		}
	}
}
