package main

//https://leetcode-cn.com/problems/minimum-increment-to-make-array-unique/

import "sort"
import "fmt"

func main() {
    var a = [] int{3, 2, 1, 2, 1, 7}
    res := minIncrementForUnique(a)
    fmt.Println(res)
}

func minIncrementForUnique(A []int) int {
    sort.Ints(A)


    temp := 0
    for i := 1; i < len(A); i++ {
        if (A[i] <= A[i-1]) {
            temp = A[i-1] - A[i] + 1 + temp
            A[i] = A[i-1] + 1
        }
    }

    return temp
}
