package main

import "fmt"

func removeElement(nums []int, val int) int {
	k := 0
	for i := 0; i < len(nums) - 1; i++ {
		if nums[i]!= val{
			nums[k] = nums[i]
			k++

		}
	}
	return k
}

func main()  {
	intlist := []int{3,2,2,3}

	fmt.Println(removeElement(intlist,3))
}