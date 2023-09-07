package list

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
