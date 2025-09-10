package main

import "fmt"

type Node[T any] struct {
	value T
	next  *Node[T]
}

type Queue[T any] struct {
	head *Node[T]
	tail *Node[T]
	size int
}

func newQueue[T any](values []T) *Queue[T] {

	if len(values) == 0 {
		return &Queue[T]{}
	}

	q := &Queue[T]{
		head: nil,
		tail: nil,
		size: 0,
	}

	for _, val := range values {
		n := &Node[T]{
			value: val,
			next:  nil,
		}

		if q.head == nil {
			q.head = n
			q.tail = n
			q.size++
			continue
		}

		if q.tail != nil {
			q.tail.next = n
			n.next = nil
			q.tail = n
			q.size++
		}
	}
	return q
}

func (q *Queue[T]) Enqueue(value T) {
	n := &Node[T]{
		value: value,
		next:  nil,
	}

	if q.head == nil {
		q.head = n
		q.tail = n
		q.size++
		return
	}

	q.tail.next = n
	q.tail = n
	q.size++
}

func (q *Queue[T]) DeQueue() (T, bool) {

	var zero T

	if q.head == nil {
		return zero, false
	}

	value := q.head.value
	q.head = q.head.next

	if q.head == nil {
		q.tail = nil
	}

	q.size--

	return value, true
}

func main() {
	fmt.Println("=== Testing Queue with ints ===")
	q1 := newQueue([]int{10, 20, 30})

	q1.Enqueue(40)
	q1.Enqueue(50)

	for i := 0; i < 6; i++ {
		if val, ok := q1.DeQueue(); ok {
			fmt.Printf("Dequeued: %d (size=%d)\n", val, q1.size)
		} else {
			fmt.Println("Queue empty!")
		}
	}

	fmt.Println("\n=== Testing Queue with strings ===")
	q2 := newQueue([]string{"apple", "banana"})
	q2.Enqueue("cherry")

	for i := 0; i < 4; i++ {
		if val, ok := q2.DeQueue(); ok {
			fmt.Printf("Dequeued: '%s'\n", val)
		} else {
			fmt.Println("Queue empty!")
		}
	}

	fmt.Println("\n=== Testing empty queue ===")
	q3 := newQueue([]int{})
	if val, ok := q3.DeQueue(); !ok {
		fmt.Printf("Dequeue on empty: ok=%v, value=%v\n", ok, val)
	}
	q3.Enqueue(100)
	if val, ok := q3.DeQueue(); ok {
		fmt.Printf("Dequeued: %v\n", val)
	}
}
