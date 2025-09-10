package main

import (
	"LeetCodeSolutions/linked_lists"
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
	node1 := &linked_lists.ListNode{Val: 1}
	node2 := &linked_lists.ListNode{Val: 2}
	node3 := &linked_lists.ListNode{Val: 3}
	node1.Next = node2
	node2.Next = node3

	node3.Next = node2
	fmt.Println(linked_lists.HasCycle(node1))
	fmt.Println(linked_lists.HasCycle2(node1))
}
