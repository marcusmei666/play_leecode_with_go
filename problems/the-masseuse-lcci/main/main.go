package main

//https://leetcode-cn.com/problems/the-masseuse-lcci/submissions/

//https://leetcode-cn.com/problems/house-robber/solution/dong-tai-gui-hua-jie-ti-si-bu-zou-xiang-jie-cjavap/

func massage(nums []int) int {
    k1, k2, temp := 0, 0, 0;
    for _, n := range nums {
        temp = getMax(k2+n, k1)
        k2 = k1
        k1 = temp
    }
    return temp
}

func getMax(a, b int) int {
    if a > b {
        return a
    }
    return b
}
