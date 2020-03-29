package main

//https://leetcode-cn.com/problems/x-of-a-kind-in-a-deck-of-cards/

func hasGroupsSizeX(deck []int) bool {
    if len(deck) == 1 {
        return false
    }
    var temp = make(map[int]int)
    for _, v := range deck {
        temp[v]++
    }
    g := -1
    for _, v := range temp {
        if v > 0 {
            if g == -1 {
                g = v
            } else {
                g = gcd(v, g)
            }
        }
    }
    return g >= 2
}

func gcd(x, y int) int {
    if x == 0 {
        return y
    } else {
        return gcd(y%x, x)
    }
}
