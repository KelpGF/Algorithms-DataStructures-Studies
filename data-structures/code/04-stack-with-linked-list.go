package main

import (
	"errors"
	"fmt"
)

type Node struct {
	value string
	next  *Node
}

func NewNode(value string, next *Node) *Node {
	return &Node{
		value: value,
		next:  next,
	}
}

type Stack struct {
	head *Node
	size int
}

func NewStack() *Stack {
	return &Stack{
		size: 0,
	}
}

func (s *Stack) Size() int {
	return s.size
}

func (s *Stack) IsEmpty() bool {
	return s.Size() == 0
}

func (s *Stack) Push(value string) {
	if s.IsEmpty() {
		s.head = NewNode(value, nil)
		s.size++

		return
	}

	newNode := NewNode(value, s.head)
	s.head = newNode
	s.size++
}

func (s *Stack) Pop() (string, error) {
	if s.IsEmpty() {
		return "", errors.New("The Stack is empty")
	}

	oldHead := s.head
	data := oldHead.value
	s.head = oldHead.next
	s.size--

	oldHead.value = ""
	oldHead.next = nil

	return data, nil
}
func (s *Stack) Peek() (string, error) {
	if s.IsEmpty() {
		return "", errors.New("The Stack is empty")
	}

	return s.head.value, nil
}

func main() {
	s := NewStack()

	s.Push("aaaa")
	fmt.Println(s.Size())
	s.Push("bbbbb")
	fmt.Println(s.Size())
	s.Push("ccccc")
	fmt.Println(s.Size())

	val, _ := s.Pop()
	fmt.Println(val)
	fmt.Println(s.Size())

	val, _ = s.Pop()
	fmt.Println(val)
	fmt.Println(s.Size())

	val, _ = s.Pop()
	fmt.Println(val)
	fmt.Println(s.Size())

	val, err := s.Pop()
	fmt.Println(val, err)
	fmt.Println(s.Size())
}
