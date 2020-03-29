package main

import "fmt"

//https://leetcode-cn.com/problems/diagonal-traverse/
func findDiagonalOrder(matrix [][]int) []int {
	sum := len(matrix) + len(matrix[0])
	a := make([]int,0)
	for i := 0; i <= sum; i++ {
		if i%2 == 0{
			j:= i
			if j > len(matrix){
				j = len(matrix)
			}
			for ; j >= 0; j--{
				a = append(a,matrix[j][i-j])
				if (i - j) > len(matrix[0]){
					break
				}
			}
		}else {
			j:= i
			if j > len(matrix[0]){
				j = len(matrix[0])
			}
			for j := i; j >= 0; j--{
				a = append(a,matrix[i-j][j])
				if j > len(matrix){
					break
				}
			}
		}
	}
	return a
}

func main()  {
	fmt.Println(2%2)
}