package main

//https://leetcode-cn.com/problems/as-far-from-land-as-possible/

type points struct {
    x, y int
}

func maxDistance(grid [][]int) int {
    gn := len(grid)
    if gn == 0 {
        return 0
    }

    var queue []*points

    for i := 0; i < gn; i++ {
        for j := 0; j < gn; j++ {
            if grid[i][j] == 1 {
                queue = append(queue, &points{x: i, y: j})
            }
        }
    }
    if len(queue) == gn*gn || len(queue) == 0 {
        return -1
    }
    var cut int
    for len(queue) != 0 {
        cut ++
        temp := queue
        queue = []*points{}

        for _, v := range temp {
            i := v.x
            j := v.y
            if i-1 >= 0 && grid[i-1][j] == 0 {
                grid[i-1][j] = 2
                queue = append(queue, &points{x: i - 1, y: j})
            }
            if i+1 < gn && grid[i+1][j] == 0 {
                grid[i+1][j] = 2
                queue = append(queue, &points{x: i + 1, y: j})
            }
            if j-1 >= 0 && grid[i][j-1] == 0 {
                grid[i][j-1] = 2
                queue = append(queue, &points{x: i, y: j - 1})
            }
            if j+1 < gn && grid[i][j+1] == 0 {
                grid[i][j+1] = 2
                queue = append(queue, &points{x: i, y: j + 1})
            }
        }

    }
    return cut-1
}
