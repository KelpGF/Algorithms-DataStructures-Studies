package main

import "errors"

type Node struct {
	value string
	prev  *Node
	next  *Node
}

func NewNode(value string, prev, next *Node) *Node {
	return &Node{
		value: value,
		prev:  prev,
		next:  next,
	}
}

type DoubleLinkedList struct {
	head *Node
	tail *Node
	size int
}

func (dl *DoubleLinkedList) Clear() {
	trav := dl.head

	for trav != nil {
		next := trav.next

		trav.prev = nil
		trav.next = nil

		trav = next
	}

	dl.head = nil
	dl.tail = nil
	dl.size = 0
}

func (dl *DoubleLinkedList) Size() int {
	return dl.size
}

func (dl *DoubleLinkedList) IsEmpty() bool {
	return dl.Size() == 0
}

func (dl *DoubleLinkedList) incSize() {
	dl.size++
}

func (dl *DoubleLinkedList) decSize() {
	dl.size--
}

func (dl *DoubleLinkedList) Add(item string) {
	dl.AddLast(item)
}

func (dl *DoubleLinkedList) AddFirst(item string) {
	if dl.IsEmpty() {
		node := NewNode(item, nil, nil)
		dl.head = node
		dl.tail = node
	} else {
		node := NewNode(item, nil, dl.head)
		dl.head.prev = node
		dl.head = node
	}

	dl.incSize()
}

func (dl *DoubleLinkedList) AddLast(item string) {
	if dl.IsEmpty() {
		node := NewNode(item, nil, nil)
		dl.head = node
		dl.tail = node
	} else {
		node := NewNode(item, dl.tail, nil)
		dl.tail.next = node
		dl.tail = node
	}

	dl.incSize()
}

func (dl *DoubleLinkedList) PeekFirst() (string, error) {
	if dl.IsEmpty() {
		return "", errors.New("The list is empty")
	}

	return dl.head.value, nil
}

func (dl *DoubleLinkedList) PeekLast() (string, error) {
	if dl.IsEmpty() {
		return "", errors.New("The list is empty")
	}

	return dl.tail.value, nil
}

func (dl *DoubleLinkedList) RemoveFirst() (string, error) {
	if dl.IsEmpty() {
		return "", errors.New("The list is empty")
	}

	data := dl.head.value
	dl.head = dl.head.next
	dl.decSize()

	if dl.IsEmpty() {
		dl.tail = nil
	} else {
		dl.head.prev = nil
	}

	return data, nil
}

func (dl *DoubleLinkedList) RemoveLast() (string, error) {
	if dl.IsEmpty() {
		return "", errors.New("The list is empty")
	}

	data := dl.tail.value
	dl.tail = dl.tail.prev
	dl.decSize()

	if dl.IsEmpty() {
		dl.head = nil
	} else {
		dl.tail.next = nil
	}

	return data, nil
}

func (dl *DoubleLinkedList) Remove(node *Node) string {
	if node.prev == nil {
		data, _ := dl.RemoveFirst()
		return data
	}
	if node.next == nil {
		data, _ := dl.RemoveLast()
		return data
	}

	node.prev.next = node.next
	node.next.prev = node.prev

	data := node.value

	node.value = ""
	node.prev = nil
	node.next = nil

	dl.decSize()

	return data
}
