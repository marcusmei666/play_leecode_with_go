package main

import (
    "fmt"
)

//https://leetcode-cn.com/problems/shu-zu-zhong-zhong-fu-de-shu-zi-lcof/
func main()  {
    a := []int{2, 1, 3, 2}
    fmt.Println(a)
    res := findRepeatNumber(a)
    fmt.Println(res)
}

func findRepeatNumber(nums []int) int {
    for i:=0;i<len(nums);i++{
        if i == nums[i]{
            continue
        }else if nums[i] == nums[nums[i]]{
            return nums[i]
        }else{
            nums[i], nums[nums[i]] = nums[nums[i]], nums[i]
        }
    }
    return -1

}

