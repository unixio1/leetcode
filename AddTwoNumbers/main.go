package main

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

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	result := &ListNode{}
	previous := result
	nodeOne := l1
	nodeTwo := l2
	remaining := 0
	for nodeOne != nil || nodeTwo != nil {
		sum := remaining
		sum += read(&nodeOne)
		sum += read(&nodeTwo)
		if sum >= 10 {
			remaining = 1
			sum = sum % 10
		} else {
			remaining = 0
		}
		previous.Val = sum
		if nodeOne == nil && nodeTwo == nil {
			break
		}
		nextNode := &ListNode{}
		previous.Next = nextNode
		previous = nextNode

	}
	if remaining > 0 {
		previous.Next = &ListNode{
			Val:  remaining,
			Next: nil,
		}
	}
	return result
}

func read(l **ListNode) int {
	if *l == nil {
		return 0
	}
	val := (*l).Val
	*l = (*l).Next
	return val
}
