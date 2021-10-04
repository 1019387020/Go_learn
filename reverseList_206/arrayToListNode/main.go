package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func ArrayToList(nums []int, head *ListNode) *ListNode {
	fnode := head
	for _, num := range nums {
		temp := ListNode{Val: num}
		head.Next = &temp
		head = &temp
	}
	return fnode.Next
}

