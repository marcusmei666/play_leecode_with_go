package main

//https://leetcode-cn.com/problems/er-wei-shu-zu-zhong-de-cha-zhao-lcof/

func findNumberIn2DArray(matrix [][]int, target int) bool {
    if len(matrix) == 0 || len(matrix[0]) == 0 {
        return false
    }
    col ,row := 0,len(matrix[0]) -1
    for row >= 0 && col < len(matrix) {
        if matrix[col][row] == target {
            return true
        }else if matrix[col][row] > target {
            row --
        }else {
            col ++
        }
    }

    return false
}