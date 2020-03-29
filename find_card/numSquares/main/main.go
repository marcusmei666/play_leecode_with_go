package main

import "fmt"

func numSquares(n int) int {
	var dp []int
	dp = make([]int,n)
	for i := 1; i <= n; i++ {
		dp[i] = i // 最坏的情况就是每次+1
		for j := 1; i - j * j >= 0; j++ {
			dp[i] = Min(dp[i], dp[i - j * j] + 1) // 动态转移方程
		}
	}

 	return dp[n]
}

func Min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func main()  {
	fmt.Println(numSquares(12))
}