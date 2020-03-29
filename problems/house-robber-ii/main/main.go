package main

//https://leetcode-cn.com/problems/house-robber-ii/
//https://leetcode-cn.com/problems/house-robber-ii/solution/tong-yong-si-lu-tuan-mie-da-jia-jie-she-wen-ti-by-/
func rob(nums []int) int {
    l := len(nums)
    if l == 0 {
        return 0
    } else if l == 1 {
        return nums[0]
    } else if l == 2 {
        return getMax(nums[0], nums[1])
    }
    return getMax(massage(nums[1:]), massage(nums[:len(nums)-1]))
}
func getMax(a, b int) int {
    if a > b {
        return a
    }
    return b
}
func massage(nums []int) int {
    k1, k2, temp := 0, 0, 0;
    for _, n := range nums {
        temp = getMax(k2+n, k1)
        k2 = k1
        k1 = temp
    }
    return temp
}
