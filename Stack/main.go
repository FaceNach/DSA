package main

import "fmt"

type Node[T any] struct {
	next  *Node[T]
	value T
}

type Stack[T any] struct {
	head *Node[T]
	size int
}

func NewStack[T any](vals []T) *Stack[T] {
	s := &Stack[T]{}

	for i := len(vals) - 1; i >= 0; i-- {
		s.push(vals[i])
	}

	return s
}

func (s *Stack[T]) push(val T) {
	n := &Node[T]{
		next:  s.head,
		value: val,
	}

	s.head = n
	s.size++
}

func (s *Stack[T]) peek() T {

	var zero T

	if s.head == nil {
		return zero
	}

	return s.head.value
}

func (s *Stack[T]) pop() T {

	var zero T

	if s.head == nil {
		return zero
	}

	value := s.head.value
	s.head = s.head.next
	s.size--

	return value

}

func (s *Stack[T]) isEmpty() bool {
	return s.head == nil
}

func (s *Stack[T]) getSize() int {
	return s.size
}

func main() {
	stack := NewStack([]int{1, 2, 3, 4, 5})

	fmt.Println("Peek:", stack.peek())
	fmt.Println("Size:", stack.getSize())

	popped := stack.pop()
	fmt.Println("Popped:", popped)

	fmt.Println("New peek:", stack.peek())
	fmt.Println("Size:", stack.getSize())

	// Test with empty stack
	emptyStack := NewStack[string]([]string{})
	fmt.Println("Empty peek:", emptyStack.peek())
	fmt.Println("Empty pop:", emptyStack.pop())
}
