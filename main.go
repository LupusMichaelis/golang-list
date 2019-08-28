package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args

	var list = List{nil}
	for _, arg := range args {
		list.Add(arg)
	}

	list.Print()
	list.Reverse()
	list.Print()

	return
}

type List struct {
	first *Node
}

type Node struct {
	next    *Node
	payload string
}

func (list *List) Add(payload string) {
	var newNode = &Node{payload: payload}

	if list.first != nil {
		newNode.next = list.first
	}

	list.first = newNode
}

func (list *List) Print() {
	fmt.Printf("List [")

	if list.first != nil {
		list.first.Print()
	}

	fmt.Printf("]\n")
}

func (list *List) Reverse() {
	var first, second, third *Node

	first = list.first

	if first == nil {
		return
	}

	for first.next != nil {

		second = first.next
		if second == nil {
			return
		}

		third = second.next

		first.next = third
		second.next = list.first

		list.first = second
	}
}

func (node *Node) Print() {
	fmt.Printf("'%s' ", node.payload)

	if node.next != nil {
		node.next.Print()
	}
}
