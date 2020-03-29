package main

import (
	"fmt"
	"sort"
)

func arrayPairSum(nums []int) (sum int) {
	sort.Ints(nums)
	for i := 0; i < len(nums); i+=2 {
		sum += nums[i]
	}
	return
}

func main()  {

	intList := [] int {2, 4, -3, 5, 7, 6, 9, 8, 1, 0}

	fmt.Println(arrayPairSum(intList))

	fmt.Println(intList)
}