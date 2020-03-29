package main

// https://leetcode-cn.com/problems/cong-wei-dao-tou-da-yin-lian-biao-lcof/

//Definition for singly-linked list.
type ListNode struct {
    Val  int
    Next *ListNode
}

func reversePrint(head *ListNode) []int {
    if head == nil {
        return nil
    }
    var temp []int
    for head != nil {
        temp = append(temp, head.Val)
        head = head.Next
    }

    for i, j := 0, len(temp)-1; i < j; {
        temp[i], temp[j] = temp[j], temp[i]
        i++
        j--
    }
    return temp
}
