package main

//https://leetcode-cn.com/problems/binary-search/
func search(nums []int, target int) int {
    var len = len(nums) - 1
    index,l,r := -1,0,len
    if len == -1 {
       return index
    }
    for l <= r  {
        mid := l + ((r - l) >> 2)
        if nums[mid] == target{
            return mid
        }else if nums[mid] > target{
            r = mid - 1
        }else {
            l = mid + 1
        }
    }
    return index
}

func main()  {

}