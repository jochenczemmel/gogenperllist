package main

import "math/big" 

// IntList contains a slice of big.Int values
type IntList struct {
	elements []big.Int
}

// NewIntList creates a new List of big.Int values
// with optional initial values
func NewIntList(values ...big.Int) *IntList {
	return &IntList{values}
}

// Len returns the length of the underlying slice
func (l *IntList) Len() int {
	return len(l.elements)
}

// Cap returns the capacity of the underlying slice
func (l *IntList) Cap() int {
	return cap(l.elements)
}

// Push appends an element at the end of the list
func (l *IntList) Push(element big.Int) {
	l.elements = append(l.elements, element)
}

// Unshift adds an element at the beginning of the list
func (l *IntList) Unshift(element big.Int) {
	newl := append([]big.Int{element}, l.elements...)
	l.elements = newl
}

// Shift returns the first element and removes it from the list
// returns zero value if list is empty
func (l *IntList) Shift() big.Int {
	if len(l.elements) == 0 {
		return *new(big.Int)
	}
	first := l.elements[0]
	l.elements = l.elements[1:]
	return first
}

// Pop returns the last element and removes it from the list
// returns zero value if list is empty
func (l *IntList) Pop() big.Int {
	if len(l.elements) == 0 {
		return *new(big.Int)
	}
	lastIndex := len(l.elements) - 1
	last := l.elements[lastIndex]
	l.elements = l.elements[:lastIndex]
	return last
}

// Foreach iterates over the slice by calling the given function
// providing the slice index and the value
func (l *IntList) Foreach(f func(int, big.Int)) {
	for index, value := range l.elements {
		f(index, value)
	}
}

// PopChecked returns the last element and removes it from the list
// returns the value and true if the list contains values,
// else the zero value and false
func (l *IntList) PopChecked() (big.Int, bool) {
	if len(l.elements) == 0 {
		r := new(big.Int)
		return *r, false
	}
	return l.Pop(), true
}

// ShiftChecked returns the first element and removes it from the list
// returns the value and true if the list contains values,
// else the zero value and false
func (l *IntList) ShiftChecked() (big.Int, bool) {
	if len(l.elements) == 0 {
		r := new(big.Int)
		return *r, false
	}
	return l.Shift(), true
}

// ElementAt returns the element at the given position.
// It panics if index is invalid (default slice behaviour)
func (l *IntList) ElementAt(index int) big.Int {
	return l.elements[index]
}

// Grep iterates over the slice by calling the given function
// providing the slice index and the value
// and returning a new List that contains all element where
// the given function returned true
func (l *IntList) Grep(f func(int, big.Int) bool) IntList {
	result := IntList{}
	for index, value := range l.elements {
		if f(index, value) {
			result.elements = append(result.elements, value)
		}
	}
	return result
}

// Map iterates over the slice by calling the given function
// providing the slice index and the value
// and returning a new List that contains all element
// with the value returned by the function
func (l *IntList) Map(f func(int, big.Int) big.Int) IntList {
	result := IntList{}
	for index, value := range l.elements {
		result.elements = append(result.elements, f(index, value))
	}
	return result
}

// Any iterates over the slice by calling the given function
// providing the slice index and the value
// and returning true if one of the function results is true
// else false
func (l *IntList) Any(f func(int, big.Int) bool) bool {
	for index, value := range l.elements {
		if f(index, value) == true {
			return true
		}
	}
	return false
}

// All iterates over the slice by calling the given function
// providing the slice index and the value
// and returning true if alle of the function results are true
// else false
func (l *IntList) All(f func(int, big.Int) bool) bool {
	for index, value := range l.elements {
		if f(index, value) == false {
			return false
		}
	}
	return true
}

// Reduce iterates over the slice by calling the given function
// providing the cumulated value, the slice index, and the value
func (l *IntList) Reduce(f func(big.Int, int, big.Int) big.Int) big.Int {
	var cum big.Int
	for index, value := range l.elements {
		if index == 0 {
			cum = value
		} else {
			cum = f(cum, index, value)
		}
	}
	return cum
}

