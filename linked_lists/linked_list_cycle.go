package linked_lists

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// решение при помощи мапы
func HasCycle(head *ListNode) bool {
	if head == nil {
		return false
	}
	visitedNodes := make(map[*ListNode]bool)
	for head.Next != nil {
		if ok := visitedNodes[head]; ok {
			return true
		} else {
			visitedNodes[head] = true
			head = head.Next
		}
	}
	return false
}

// алгоритм заяц и черепаха
func HasCycle2(head *ListNode) bool {
	if head == nil {
		return false
	}
	p1 := head
	p2 := head

	for p2 != nil && p2.Next != nil {
		p1 = p1.Next
		p2 = p2.Next.Next
		if p1 == p2 {
			return true
		}
	}
	return false
}
