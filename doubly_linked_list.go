// Write your answer here, and then test your code.
// Your job is to implement the Add(index int, v T) method.

package dl_list

import (
	// "errors"
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
	if index == 0 {
		newNode := &Node[T]{value: v, next: nil, prev: nil}
		if l.size == 0 {
			l.head = newNode
			l.tail = newNode
		}
		l.size=1
	}
	
	return nil
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