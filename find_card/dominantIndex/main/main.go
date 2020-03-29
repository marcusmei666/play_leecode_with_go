package main

import "fmt"

func dominantIndex1(nums []int) int {
	if len(nums) == 1{
		return 0
	}
	var max, index int
	var flag bool
	for i:= 0; i < len(nums); i++{
		if nums[i] > max{
			max = nums[i]
			index = i
		}
	}
	fmt.Println(max,index,flag)
	for i := 0; i < len(nums); i++ {
		if i != index{
			if nums[i]*2 >= max{
				flag = true
			}
		}
	}
	fmt.Println(max,index,flag)
	if flag{
		return -1
	}else {
		return index
	}
}

func dominantIndex(nums []int) int {
	if len(nums) == 1{
		return 0
	}
	var max, index,sec int

	for i:= 0; i < len(nums); i++{
		if nums[i] > sec{
			sec = nums[i]
			if sec > max{
				sec, max = max,sec
				index = i
			}
		}
	}

	if sec*2 > max{
		return -1
	}else {
		return index
	}
}

func main()  {
	var a = []int{0,0,0,1}
	fmt.Println(dominantIndex1(a))


}