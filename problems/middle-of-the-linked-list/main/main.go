package main

//https://leetcode-cn.com/problems/middle-of-the-linked-list/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
    Val  int
    Next *ListNode
}

func middleNode(head *ListNode) *ListNode {
    if head == nil || head.Next == nil{
        return head
    }

    fast, slow := head, head

    for fast.Next != nil && fast.Next.Next != nil{
        fast = fast.Next.Next
        slow = slow.Next
    }
    if fast.Next != nil {
        return slow.Next
    }else {
        return  slow
    }
}
