package main

import (
	"fmt"
	"strconv"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(nums []int) *TreeNode {
	if len(nums) == 0 || nums[0] == -1 {
		return nil
	}
	nodes := make([]*TreeNode, len(nums))
	nodes[0] = &TreeNode{Val: nums[0]}
	for i := 1; i < len(nums); i++ {
		if nums[i] < 0 {
			continue
		}

		node := &TreeNode{Val: nums[i]}
		father := nodes[(i-1)/2]
		if father == nil {
			continue
		}
		if i%2 == 1 {
			father.Left = node
		}
		if i%2 == 0 {
			father.Right = node
		}

		nodes[i] = node
	}
	return nodes[0]
}

func isLeaf(node *TreeNode) bool {
	return node.Left == nil && node.Right == nil
}

func preOrder(root *TreeNode) []*TreeNode {
	visit := []*TreeNode{}
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		visit = append(visit, queue[0])
		if queue[0].Left != nil {
			queue = append(queue, queue[0].Left)
		}
		if queue[0].Right != nil {
			queue = append(queue, queue[0].Right)
		}
		queue = queue[1:]
	}
	return visit
}

func inorder(root *TreeNode) []*TreeNode {
	visit := []*TreeNode{}
	stack := []*TreeNode{}
	cur := root
	for cur != nil || len(stack) > 0 {
		for ; cur != nil; cur = cur.Left {
			stack = append(stack, cur)
		}
		cur = stack[len(stack)-1]
		visit = append(visit, cur)
		stack = stack[:len(stack)-1]
		cur = cur.Right
	}
	return visit
}

func printTreeNodeList(list []*TreeNode) {
	nums := []int{}
	for i := 0; i < len(list); i++ {
		nums = append(nums, list[i].Val)
	}
	fmt.Println(nums)
}

func printTree(root *TreeNode) {
	nums := []string{}
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		l := len(queue)
		for i := 0; i < l; i++ {
			if queue[i] == nil {
				nums = append(nums, "null")
				continue
			}
			nums = append(nums, strconv.Itoa(queue[i].Val))
			if isLeaf(queue[i]) {
				continue
			}
			queue = append(queue, queue[i].Left)
			queue = append(queue, queue[i].Right)
		}
		queue = queue[l:]
	}
	fmt.Println(nums)
}

// func main() {
// nums := []int{1, 3, -1, -1, 2}
// root := buildTree(nums)
//
// printTreeNodeList(preOrder(root))
// }
