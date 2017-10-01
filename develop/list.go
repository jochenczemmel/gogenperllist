/*
implements some perl list functions for MYTYPE lists
Base for template for generate tool
*/
package main

type MYTYPE = int

type MYTYPEList struct {
	elements []MYTYPE
}

// creates a new list with optional initial values
func NewMYTYPELIST(values ...MYTYPE) *MYTYPEList {
	return &MYTYPEList{values}
}

// return len of list
func (l *MYTYPEList) Len() int {
	return len(l.elements)
}

// return cap of list
func (l *MYTYPEList) Cap() int {
	return cap(l.elements)
}

// append element to list
func (l *MYTYPEList) Push(element MYTYPE) {
	l.elements = append(l.elements, element)
}

// prepend element to list
func (l *MYTYPEList) Unshift(element MYTYPE) {
	newl := append([]MYTYPE{element}, l.elements...)
	l.elements = newl
}

// get first element of list
// or zero value if list is empty
func (l *MYTYPEList) Shift() MYTYPE {
	if len(l.elements) == 0 {
		r := new(MYTYPE)
		return *r
	}
	first := l.elements[0]
	l.elements = l.elements[1:]
	return first
}

// get last element of list
// or zero value if list is empty
func (l *MYTYPEList) Pop() MYTYPE {
	if len(l.elements) == 0 {
		r := new(MYTYPE)
		return *r
	}
	lastIndex := len(l.elements) - 1
	last := l.elements[lastIndex]
	l.elements = l.elements[:lastIndex]
	return last
}

// iterate over list elements
func (l *MYTYPEList) Foreach(f func(int, MYTYPE)) {
	for index, value := range l.elements {
		f(index, value)
	}
}

// get last element of list
// true if list is not empty, else false
func (l *MYTYPEList) PopChecked() (MYTYPE, bool) {
	if len(l.elements) == 0 {
		r := new(MYTYPE)
		return *r, false
	}
	return l.Pop(), true
}

// get first element of list
// true if list is not empty, else false
func (l *MYTYPEList) ShiftChecked() (MYTYPE, bool) {
	if len(l.elements) == 0 {
		r := new(MYTYPE)
		return *r, false
	}
	return l.Shift(), true
}

// get element at index
// no error handling, default slice behaviour!
func (l *MYTYPEList) ElementAt(index int) MYTYPE {
	return l.elements[index]
}
