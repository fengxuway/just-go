package main

import "fmt"

type Object interface{}
type Node struct {
	data Object
	next *Node
}

type Link struct {
	head   *Node
	length int
	last   *Node
}

func (link *Link) pop() Object {
	if link.length > 0 {
		value := link.head.data
		link.head = link.head.next
		link.length--
		return value
	}
	return nil
}

func (link *Link) push(value Object) {
	node := &Node{data: value, next: nil}
	if link.last != nil {
		link.last.next = node
		link.last = node
	} else {
		link.head = node
		link.last = node
	}
	link.length = link.length + 1
}

func (link *Link) loop() {
	node := link.head
	for {
		if node == nil {
			break
		}
		fmt.Println(node.data)
		node = node.next
	}
}

func main() {
	//n := &Node{data: 1, next: nil}
	link := Link{head: nil, last: nil, length: 0}
	link.push(1)
	link.push(4)
	link.push(6)
	link.loop()
	fmt.Println("=======")
	fmt.Print(link.pop())
	fmt.Println("=======")
	link.loop()
	fmt.Println(link.length)
}
