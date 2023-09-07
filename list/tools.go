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

// ReverseNodesValuesBetween helps to swap near right and left nodes values between their node.
func ReverseNodesValuesBetween(head *Node, left int, right int) *Node {
	if head == nil || left == right {
		return head
	}

	dummy := &Node{0, head}
	prev := dummy

	for i := 0; i < left-1; i++ {
		prev = prev.Next
	}

	current := prev.Next

	for i := 0; i < right-left; i++ {
		nextNode := current.Next
		current.Next = nextNode.Next
		nextNode.Next = prev.Next
		prev.Next = nextNode
	}

	return dummy.Next
}
