package main

import (
	"fmt"
)

type List struct {
	first *Node
}

type Node struct {
	next    *Node
	payload interface{}
}

// XXX find a way to provide good interface for any type of payload
// XXX see issue with https://github.com/golang/go/wiki/InterfaceSlice
func (list *List) AddMany(payloadList ...string) {
	for _, payload := range payloadList {
		list.Add(payload)
	}
}

func (list *List) Add(payload interface{}) {
	var newNode = &Node{payload: payload}

	if list.first != nil {
		newNode.next = list.first
	}

	list.first = newNode
}

func (list List) String() (output string) {
	output = fmt.Sprintf("List [")

	if list.first != nil {
		output += list.first.String()
	}

	output += fmt.Sprintf("]")

	return
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

func (node Node) String() (output string) {
	output = fmt.Sprintf("'%v'", node.payload)

	if node.next != nil {
		output += ","
		output += node.next.String()
	}

	return
}
