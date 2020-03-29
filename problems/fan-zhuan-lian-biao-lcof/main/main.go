package main

//https://leetcode-cn.com/problems/fan-zhuan-lian-biao-lcof/

//Definition for singly-linked list.
type ListNode struct {
    Val  int
    Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
    if head == nil {
        return nil
    }
    cur := head
    for head.Next != nil  {
        temp := head.Next.Next
        head.Next.Next = cur
        cur = head.Next
        head.Next = temp
    }
    return cur
}
