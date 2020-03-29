package main

//https://leetcode-cn.com/problems/climbing-stairs/
//https://leetcode-cn.com/problems/climbing-stairs/solution/pa-lou-ti-by-leetcode/
func climbStairs(n int) int {
    if n <= 2 {
        return n
    }
    dp1, dp2 := 1, 2
    for i := 3; i <= dp2; i++ {
        dp1, dp2 = dp2, dp1+dp2
    }
    return dp2
}
