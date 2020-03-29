package main

func mySqrt(x int) int {
    if x == 0 {
        return 0
    }
    if x <= 3 {
        return 1
    }
    l, r := 0, x/2
    for l <= r {
        mid := (l + r) / 2
        temp := mid * mid
        if temp == x {
            return mid
        } else if temp > x {
            r = mid - 1
        } else {
            l = mid + 1
        }
    }
    return 2
}
