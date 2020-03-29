package main

func rotate(nums []int, k int)  {
	k %= len(nums)
	nums = append(nums[len(nums)-k:], nums[:len(nums)-k]...)
}

