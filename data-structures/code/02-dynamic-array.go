package main

import (
	"errors"
	"fmt"
)

type DynamicArray struct {
	currentSize int
	capacity    int
	array       []string
}

func (d DynamicArray) Size() int {
	return d.currentSize
}

func (d DynamicArray) cap() int {
	return d.capacity
}

func (d DynamicArray) IsEmpty() bool {
	return d.Size() == 0
}

func (d *DynamicArray) Add(item string) {
	if d.Size() < d.cap() {
		d.array[d.Size()] = item
		d.currentSize++
		return
	}

	if d.cap() == 0 {
		d.capacity = 1
	} else {
		d.capacity *= 2
	}

	old := d.array
	d.array = make([]string, d.cap())

	for idx, item := range old {
		d.array[idx] = item
	}

	d.array[d.Size()] = item
	d.currentSize++
}

func (d *DynamicArray) RemoveAt(rIndex int) error {
	if rIndex < 0 || rIndex > d.Size()-1 {
		return errors.New("Invalid Index")
	}

	old := d.array
	d.capacity--
	d.currentSize--
	d.array = make([]string, d.cap())

	j := 0
	for i := 0; i < d.Size(); i++ {
		if i == rIndex {
			j++
		}

		d.array[i] = old[j]
		j++
	}

	return nil
}

func (d *DynamicArray) IndexOf(rItem string) int {
	for idx, item := range d.array {
		if item == rItem {
			return idx
		}
	}

	return -1
}

func (d *DynamicArray) Remove(rItem string) bool {
	idx := d.IndexOf(rItem)

	if idx == -1 {
		return false
	}

	d.RemoveAt(idx)
	return true
}

func (d *DynamicArray) Contains(rItem string) bool {
	idx := d.IndexOf(rItem)

	return idx != -1
}

func NewDynamicArray(size int) (DynamicArray, error) {
	if size < 0 {
		return DynamicArray{}, errors.New("Invalid Capacity")
	}

	return DynamicArray{
		currentSize: 0,
		capacity:    size,
		array:       make([]string, size),
	}, nil
}

func main() {
	d, _ := NewDynamicArray(0)

	fmt.Println(d.array)
	fmt.Println(cap(d.array))
	d.Add("kel")
	fmt.Println(d.array)
	fmt.Println(cap(d.array))
	d.Add("emm")
	fmt.Println(d.array)
	fmt.Println(cap(d.array))
	d.Add("jel")
	fmt.Println(d.array)
	fmt.Println(d.currentSize)
	fmt.Println(cap(d.array))

	// d.RemoveAt(0)
	d.Remove("kel")
	fmt.Println(d.array)
	fmt.Println(d.currentSize)
	fmt.Println(cap(d.array))
}
