package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func printList(l *ListNode) {
	list := []int{}
	for ; l != nil; l = l.Next {
		list = append(list, l.Val)
	}
	fmt.Println(list)
}

func buildList(list []int) *ListNode {
	head := &ListNode{}
	l := head
	for _, val := range list {
		l.Next = &ListNode{Val: val}
		l = l.Next
	}
	return head.Next
}
