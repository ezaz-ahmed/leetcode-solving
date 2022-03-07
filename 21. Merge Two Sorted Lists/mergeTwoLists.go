package main

import "fmt"

func main() {

	list1, list2 := []int{1, 2, 4}, []int{1, 3, 4}

	fmt.Println(mergeTwoLists(list1, list2))
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	start := &ListNode{
		Val:  -200,
		Next: nil,
	}
	h := start
	for list1 != nil || list2 != nil {
		if list1 == nil {
			h.Next, list2, h = list2, list2.Next, list2
			continue
		}
		if list2 == nil {
			h.Next, list1, h = list1, list1.Next, list1
			continue
		}
		if list1.Val < list2.Val {
			h.Next, list1, h = list1, list1.Next, list1
			continue
		}
		h.Next, list2, h = list2, list2.Next, list2
	}
	return start.Next
}
