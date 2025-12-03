package main

import "fmt"

type Stack[T any] struct{
	data []T
}


func newStack[T any]() *Stack[T]{
	return &Stack[T]{
		data: make([]T, 0,10),
	}
}

func newStackWithValues[T any](vals ...T) *Stack[T]{
	return &Stack[T]{
		data: vals,
	}
}

func (s *Stack[T]) push(val T){
	s.data = append(s.data, val)
}

func (s *Stack[T]) pop() (T, bool){
	if len(s.data) == 0 {
		var Zero T
		return Zero, false
	}
	
	index := len(s.data) - 1
	element := s.data[index]
	
	s.data = s.data[:index]
	
	return element, true
}

func (s *Stack[T]) peek() (T, bool){
	if len(s.data) == 0 {
		var Zero T
		return Zero, false
	}
	
	index := len(s.data) - 1
	
	
	return s.data[index], true
	
}

func main() {
	stack := newStackWithValues[int](1, 2, 3, 4, 5)
	val, _ := stack.peek()
	fmt.Printf("Peek Stack 1: %v\n", val)


	emptyStack := newStack[string]()
	emptyStack.push("Hola")
	val2, _ := emptyStack.peek()
	fmt.Printf("Peek Stack 2: %v\n", val2)
}