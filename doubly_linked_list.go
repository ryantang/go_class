// Write your answer here, and then test your code.
// Your job is to implement the Add(index int, v T) method.

package dl_list

import (
	"errors"
	"fmt"
)

// Change these boolean values to control whether you see
// the expected answer and/or hints.
const showExpectedResult = false;
const showHints = false;

type Node[T any] struct {
	value      T
	next, prev *Node[T]
}

type DoublyLinkedList[T any] struct {
	size int
	head *Node[T]
	tail *Node[T]
}

// Add appends the given value at the given index.
// Returns an error in the case that the index exceeds the list size.
func (l *DoublyLinkedList[T]) Add(index int, v T) error {

	var currentNode *Node[T]

	if l.size < index {
		return errors.New("index exceeds list size")
	}

	//head of non-empty list
	if index == 0 && l.head != nil {
		nextNode := l.head
		currentNode = &Node[T]{value: v, next: nextNode, prev: nil}
		nextNode.prev = currentNode
		l.head = currentNode
		l.size++
	}

	//empty list
	if index == 0 && l.head == nil {
		currentNode = &Node[T]{value: v, next: nil, prev: nil}
		l.head = currentNode
		l.tail = currentNode
		l.size++
	}

	//middle	
	if index > 0 && index < l.size {
		nextNode := l.atPosition(index)
		previousNode := l.atPosition(index - 1)
		currentNode = &Node[T]{value: v, next: nextNode, prev: previousNode}
		previousNode.next = currentNode
		nextNode.prev = currentNode
		l.size++
	}

	//tail
	if index > 0 && index == l.size {
		previousNode := l.atPosition(index - 1)
		currentNode = &Node[T]{value: v, next: nil, prev: previousNode}
		previousNode.next = currentNode
		l.tail = currentNode
		l.size++
	}
	return nil
}

func (l * DoublyLinkedList[T]) atPosition(index int) *Node[T] {
	if index == 0 {
		return l.head
	}

	currentNode := l.head
	for i := 0; i < index; i++ {
		currentNode = currentNode.next
	}

	return currentNode
}

func (l *DoublyLinkedList[T]) AddElements(elements []struct{
	index int
	value T
}) error {
	for _, e := range elements {
		if err := l.Add(e.index, e.value); err != nil {
			return err
		}
	}

	return nil
}

func (l *DoublyLinkedList[T]) PrintForward() string {
	if l.size == 0 {
		return ""
	}
	current := l.head
	output := "HEAD"
	for current != nil {
		output = fmt.Sprintf("%s -> %v", output, current.value)
		current = current.next
	}

	// fmt.Printf("%s <- HEAD", output)
	return fmt.Sprintf("%s -> NULL", output)
}

func (l *DoublyLinkedList[T]) PrintReverse() string {
	if l.size == 0 {
		return ""
	}
	current := l.tail
	output := "NULL"
	for current != nil {
		output = fmt.Sprintf("%s <- %v", output, current.value)
		current = current.prev
	}
	// fmt.Printf("%s <- HEAD", output)
	return fmt.Sprintf("%s <- HEAD", output)
}