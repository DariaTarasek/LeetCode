package main

import (
	"LeetCodeSolutions/binary_trees"
	"LeetCodeSolutions/linked_lists"
	"LeetCodeSolutions/strings_slices"
	_ "LeetCodeSolutions/strings_slices"
	"fmt"
)

func main() {
	//fmt.Println(maps_frequencies.TopKFrequent([]int{1, 1, 0, 2, 2, 3, 3, 3, 4, 4, 4, 4, 5, 5, 5, 5, 5}, 2))
	//slices_pointers.Merge([]int{4, 5, 6, 0, 0, 0}, 3, []int{1, 2, 3}, 3)
	//slices_pointers.MoveZeroes([]int{0, 2, 0, 3})
	//fmt.Println(slices_pointers.MaxProfit([]int{7, 1, 5, 3, 6, 4}))
	//fmt.Println(slices_pointers.MaxProfit([]int{2, 6, 1, 3, 1}))
	//fmt.Println(slices_pointers.RemoveDuplicates([]int{1}))
	//node1 := &linked_lists.ListNode{Val: 1}
	//node2 := &linked_lists.ListNode{Val: 2}
	//node3 := &linked_lists.ListNode{Val: 3}
	//node1.Next = node2
	//node2.Next = node3
	//
	//node3.Next = node2
	//fmt.Println(linked_lists.HasCycle(node1))
	//fmt.Println(linked_lists.HasCycle2(node1))
	//
	//obj := stacks_queues.Constructor()
	//obj.Push(1)
	//obj.Push(2)
	//obj.Push(5)
	//fmt.Println(obj)
	//obj.Pop()
	//fmt.Println(obj)
	//fmt.Println(obj.Peek())
	//fmt.Println(obj.Empty())

	n1 := binary_trees.TreeNode{Val: 1}
	n2 := binary_trees.TreeNode{Val: 10}
	n3 := binary_trees.TreeNode{Val: 2}
	n4 := binary_trees.TreeNode{Val: 15}
	n5 := binary_trees.TreeNode{Val: 4}
	n6 := binary_trees.TreeNode{Val: 5}
	n7 := binary_trees.TreeNode{Val: 5}

	n1.Left = &n2
	n1.Right = &n3
	n2.Left = &n4
	n2.Right = &n6
	n3.Left = &n5
	n3.Right = &n7

	//fmt.Println(binary_trees.MaxDepth(&n1))
	//fmt.Println(binary_trees.IsSymmetric(&n1))
	//fmt.Println(binary_trees.InvertTree(&n1))

	// fmt.Println(other.FizzBuzz(15))

	//fmt.Println(other.ClimbStairs(1))

	//fmt.Println(other.Search([]int{-5, -1, 0, 3, 5, 9, 12, 18}, 18))

	//slices_pointers.RemoveElement([]int{0, 1, 2, 2, 3, 0, 4, 2}, 2)

	//[5,4,3,2,1,1]
	listN1 := linked_lists.ListNode{Val: 5}
	listN2 := linked_lists.ListNode{Val: 4}
	listN3 := linked_lists.ListNode{Val: 3}
	listN4 := linked_lists.ListNode{Val: 2}
	listN5 := linked_lists.ListNode{Val: 1}

	listN1.Next = &listN2
	listN2.Next = &listN3
	listN3.Next = &listN4
	listN4.Next = &listN5

	//printLinkedList(&listN1)
	//l := linked_lists.RemoveElements(&listN1, 1)
	//printLinkedList(l)

	//fmt.Println(other.SearchInsert([]int{1, 2, 3, 4, 7}, 6))
	//fmt.Println(strings_slices.ConvertToTitle(28))
	//fmt.Println(strings_slices.TitleToNumber("AAA"))

	fmt.Println(strings_slices.LengthOfLastWord("hello world   "))

}

func printLinkedList(node *linked_lists.ListNode) {
	current := node
	for current != nil {
		fmt.Print(current.Val, " -> ")
		current = current.Next
	}
	fmt.Println(nil)
}
