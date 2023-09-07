package list

import "fmt"

type Node struct {
	Value int
	Next  *Node
}

// FindLargestNode helps to find node with max value.
func FindLargestNode(head *Node) *Node {
	if head == nil {
		return nil
	}

	largest := head

	for current := head; current != nil; current = current.Next {
		if current.Value > largest.Value {
			largest = current
		}
	}

	return largest
}

// PrintLinkedList helps to print all values of linked list in console.
func PrintLinkedList(list *Node) {
	for list.Next != nil {
		fmt.Println(list.Value)
		list = list.Next
	}
}
