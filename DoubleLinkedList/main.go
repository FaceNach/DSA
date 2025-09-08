package main

import "fmt"

type Node struct {
	value int
	next  *Node
	prev  *Node
}

type LinkedList struct {
	head *Node
	tail *Node
}

func newLinkedList(values []int) *LinkedList {
	if len(values) == 0 {
		return &LinkedList{}
	}

	linkedList := &LinkedList{
		head: nil,
		tail: nil,
	}

	for _, val := range values {
		n := &Node{
			value: val,
			next:  nil,
			prev:  nil,
		}

		if linkedList.head == nil {
			linkedList.head = n
			linkedList.tail = n
			continue
		}

		if linkedList.tail != nil {
			linkedList.tail.next = n
			n.prev = linkedList.tail
			linkedList.tail = n
			linkedList.tail.next = nil
		}
	}

	return linkedList

}

func (l *LinkedList) Append(value int) {
	n := &Node{
		value: value,
		next:  nil,
		prev:  nil,
	}

	if l.head == nil {
		l.head = n
		l.tail = n
		return
	}

	l.tail.next = n
	n.prev = l.tail
	l.tail = n
	l.tail.next = nil

}

func (l *LinkedList) Insert(value, position int) {
	n := &Node{
		value: value,
		next:  nil,
		prev:  nil,
	}

	if l.head == nil {
		l.head = n
		l.tail = n
		return
	}

	length := l.numberOfNodesOnLinkedList()

	if position > length {
		return
	}

	if position == 0 {
		l.head.prev = n
		n.next = l.head
		l.head = n
		return
	}

	if length == position {
		l.Append(value)
		return
	}

	index := 0

	for i := l.head; i.next != nil; i = i.next {
		if index == position {
			i.prev.next = n
			n.prev = i.prev
			n.next = i
			i.prev = n
			return
		}

		index++
	}

}

func (l *LinkedList) RemoveValue(value int) {

	if l.head == nil {
		return
	}

	if value == l.head.value {
		l.head = l.head.next
		if l.head != nil {
			l.head.prev = nil
		} else {
			l.tail = nil
		}
		return
	}

	if value == l.tail.value {
		l.tail = l.tail.prev
		if l.tail != nil {
			l.tail.next = nil
		} else {
			l.head = nil
		}
		return
	}

	for i := l.head; i != nil; i = i.next {

		if i.value == value {
			i.prev.next = i.next
			i.next.prev = i.prev
			return
		}
	}

}

func (l *LinkedList) numberOfNodesOnLinkedList() int {
	var count int

	if l.head == nil {
		return 0
	}

	for i := l.head; ; i = i.next {
		count++
		if i.next == nil {
			break
		}
	}

	return count
}

func (l *LinkedList) printLinkedListForward() {
	if l.head == nil {
		fmt.Println("The linked list it's empty")
		return
	}

	for i := l.head; ; i = i.next {

		fmt.Printf("%v\n", i.value)

		if i.next == nil {
			break
		}
	}
}

func (l *LinkedList) findValue(value int) (bool, int) {

	if l.head == nil {
		return false, -1
	}

	if value == l.head.value {
		return true, 0
	}

	count := 1

	for i := l.head.next; i != nil; i = i.next {
		if i.value == value {
			return true, count
		}
		count++
	}

	return false, -1
}

func main() {
	values := []int{10, 20, 30, 40}
	list := newLinkedList(values)

	fmt.Println("Initial list:")
	for node := list.head; node != nil; node = node.next {
		fmt.Printf("%d ", node.value)
	}
	fmt.Println("\n---")

	list.Append(50)
	fmt.Println("After Append(50):")
	list.printLinkedListForward()
	fmt.Println("---")

	list.Insert(15, 1) // insert 15 at position 1
	fmt.Println("After Insert(15, 1):")
	list.printLinkedListForward()
	fmt.Println("---")

	list.RemoveValue(30)
	fmt.Println("After RemoveValue(30):")
	list.printLinkedListForward()
	fmt.Println("---")

	found, index := list.findValue(40)
	if found {
		fmt.Printf("Value 40 found at index %d\n", index)
	} else {
		fmt.Println("Value 40 not found")
	}

	found, index = list.findValue(100)
	if found {
		fmt.Printf("Value 100 found at index %d\n", index)
	} else {
		fmt.Println("Value 100 not found")
	}

	fmt.Printf("Number of nodes in the list: %d\n", list.numberOfNodesOnLinkedList())

}
