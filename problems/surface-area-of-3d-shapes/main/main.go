package main

//https://leetcode-cn.com/problems/surface-area-of-3d-shapes/

func surfaceArea(grid [][]int) int {
    if len(grid) == 0 {
        return 0
    }

    num, face, l := 0, 0, len(grid)
    for i := 0; i < l; i++ {
        for j := 0; j < l; j++ {
            num += grid[i][j]

            if grid[i][j] > 1 {
                face += grid[i][j] - 1
            }
            if i > 0 {
                face += getMin(grid[i][j], grid[i-1][j])
            }
            if j > 0 {
                face += getMin(grid[i][j], grid[i][j-1])
            }
        }
    }

    return num*6 - face*2
}

func getMin(a, b int) int {
    if a < b {
        return a
    }
    return b
}
