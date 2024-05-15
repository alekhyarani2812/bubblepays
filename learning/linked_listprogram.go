package learning

import (
	"fmt"

	"github.com/bubblepays/log"
)

// Insert a node at given index
type Node struct {
	data int
	next *Node
}

// function to create new node
func NewNode(val int) *Node {
	return &Node{
		data: val,
		next: nil,
	}
}
func PrintLinkedList(head *Node) {
	current := head
	for current != nil {
		log.Infof(" %d :", current.data)
		current = current.next
	}
	fmt.Println()
}
func InsertNodeAtIndexImpl(head *Node, index, data int) *Node {
	if index < 0 {
		log.Info("invalid index")
		return head
	}
	newNode := NewNode(data)
	if index == 0 {
		newNode.next = head
		return newNode
	}
	current := head
	for i := 0; i < index-1; i++ {
		current = current.next
	}
	newNode.next = current.next
	current.next = newNode
	return head
}
func InsertNodeAtIndex() {
	var head *Node
	head = NewNode(1)
	head.next = NewNode(2)
	head.next.next = NewNode(3)
	log.Info("Linked List before insertion")
	PrintLinkedList(head)
	head = InsertNodeAtIndexImpl(head, 3, 12)
	log.Info("Linked list after insertion of 12 at position 3:")
	PrintLinkedList(head)
}

func InsertLinkedList() {
	//var head *Node
	head := NewNode(15)
	head.next = NewNode(25)
	head.next.next = NewNode(33)
	head.next.next.next = NewNode(43)
	head.next.next.next.next = NewNode(56)
	log.Info("Linked List before reverse")
	PrintLinkedList(head)
	//head.LoopLinkedListMethod()
	log.Info("Linked List after reverse")
	reverseHead := ReverseLinkedList(head)
	PrintLinkedList(reverseHead)
	//reverseHead.LoopLinkedListMethod()
}
func ReverseLinkedList(head *Node) *Node {
	var prev, next *Node
	current := head
	for current != nil {
		next = current.next
		current.next = prev
		prev = current
		current = next
	}
	return prev
}
func (n *Node) LoopLinkedListMethod() {
	log.Info("********* LoopLinkedListMethod ")
	current := n
	for current != nil {
		log.Infof(" %d :", current.data)
		current = current.next
	}
	fmt.Println()
}

//initialize the nodes
// func SimpleLinkedListPrg(){
// 	var head *Node
// 	one := NewNode(1)
// 	two := NewNode(2)
// 	three := NewNode(3)

// 	one.next = two
// 	two.next =three

// 	head = one
// 	TraveseLinkedList(head)
// }
