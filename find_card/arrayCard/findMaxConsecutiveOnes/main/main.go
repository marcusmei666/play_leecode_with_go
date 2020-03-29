package main

import "fmt"

func findMaxConsecutiveOnes(nums []int) (max int) {

	index := -1
	nums = append(nums,0)
	for i, v := range nums{

		if v == 0{
			if (i - index - 1) > max{
				max = i - index - 1
			}
			index = i
		}
	}
	return
}

func main()  {
	intlist := []int{1,1,0,1,1,1}

	fmt.Println(findMaxConsecutiveOnes(intlist))
}