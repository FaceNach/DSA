package main

import "fmt"

const NodeCapacity int = 4

type UnrolledNode struct{
	elements [NodeCapacity]int
	count int
	next *UnrolledNode
}

type UnrolledLinkedList struct {
	head *UnrolledNode
}

func (ul *UnrolledLinkedList) add(val int){
	if ul.head == nil {
		ul.head = &UnrolledNode{
			elements: [NodeCapacity]int{val},
			count: 1,
			next: nil,
		}
		return 
	}
	
	current := ul.head
	for current.next != nil{
		current = current.next
	}
	
	if current.count < NodeCapacity {
		current.elements[current.count] = val
		current.count++
	}else{
		un := &UnrolledNode{
			elements: [NodeCapacity]int{val},
			count: 1,
			next: nil,
		}
		current.next = un
	}
	
}

func (ul *UnrolledLinkedList) get (index int) (int, error){
	
	var zero int
	
	if index < 0 {
		return 0, fmt.Errorf("You can't get a negative index")
	}
	
	current := ul.head
	
	for current != nil {
		if index < current.count {
			return current.elements[index], nil
		}
		index -= current.count
		current = current.next
	}
	
	return zero, fmt.Errorf("Out of index")
}


func (ul *UnrolledLinkedList) print(){
	if ul.head == nil {
		fmt.Println("Empty UnrolledLinkedList")
		return
	}
	
	current := ul.head
	
	for current != nil{
		fmt.Printf("%v ->", current.elements[:current.count])
		current = current.next
	}
}

func main() {
	list := UnrolledLinkedList{}
	
	list.add(1)
	list.add(2)
	list.add(3)
	list.add(4)
	list.add(5)
	list.add(6)
	list.add(7)
	list.add(8)
	list.add(9)
	list.add(10)
	list.add(11)
	list.add(12)
	
	get, _ := list.get(5)
	get2, err := list.get(13)
	 if err != nil {
	fmt.Println(err)		
	}else {
		fmt.Println(get2)
	}
	
	fmt.Println(get)
	list.print()
}