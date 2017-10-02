package main

// MYTYPEList contains a slice of MYTYPE values
type MYTYPEList struct {
	elements []MYTYPE
}

// NewMYTYPEList creates a new List of MYTYPE values
// with optional initial values
func NewMYTYPEList(values ...MYTYPE) *MYTYPEList {
	return &MYTYPEList{values}
}

// Len returns the length of the underlying slice
func (l *MYTYPEList) Len() int {
	return len(l.elements)
}

// Cap returns the capacity of the underlying slice
func (l *MYTYPEList) Cap() int {
	return cap(l.elements)
}

// Push appends an element at the end of the list
func (l *MYTYPEList) Push(element MYTYPE) {
	l.elements = append(l.elements, element)
}

// Unshift adds an element at the beginning of the list
func (l *MYTYPEList) Unshift(element MYTYPE) {
	newl := append([]MYTYPE{element}, l.elements...)
	l.elements = newl
}

// Shift returns the first element and removes it from the list
// returns zero value if list is empty
func (l *MYTYPEList) Shift() MYTYPE {
	if len(l.elements) == 0 {
		r := new(MYTYPE)
		return *r
	}
	first := l.elements[0]
	l.elements = l.elements[1:]
	return first
}

// Pop returns the last element and removes it from the list
// returns zero value if list is empty
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

// Foreach iterates over the slice by calling the given function
// providing the slice index and the value
func (l *MYTYPEList) Foreach(f func(int, MYTYPE)) {
	for index, value := range l.elements {
		f(index, value)
	}
}

// PopChecked returns the last element and removes it from the list
// returns the value and true if the list contains values,
// else the zero value and false
func (l *MYTYPEList) PopChecked() (MYTYPE, bool) {
	if len(l.elements) == 0 {
		r := new(MYTYPE)
		return *r, false
	}
	return l.Pop(), true
}

// ShiftChecked returns the first element and removes it from the list
// returns the value and true if the list contains values,
// else the zero value and false
func (l *MYTYPEList) ShiftChecked() (MYTYPE, bool) {
	if len(l.elements) == 0 {
		r := new(MYTYPE)
		return *r, false
	}
	return l.Shift(), true
}

// ElementAt returns the element at the given position.
// It panics if index is invalid (default slice behaviour)
func (l *MYTYPEList) ElementAt(index int) MYTYPE {
	return l.elements[index]
}
