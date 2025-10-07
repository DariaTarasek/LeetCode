package main

import (
	"LeetCodeSolutions/linked_lists"
	_ "LeetCodeSolutions/strings_slices"
	"fmt"
)

func main() {
}

func printLinkedList(node *linked_lists.ListNode) {
	current := node
	for current != nil {
		fmt.Print(current.Val, " -> ")
		current = current.Next
	}
	fmt.Println(nil)
}
