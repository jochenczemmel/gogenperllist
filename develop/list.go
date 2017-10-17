package main

// REPlACEIMPORT

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
		return *new(MYTYPE)
	}
	first := l.elements[0]
	l.elements = l.elements[1:]
	return first
}

// Pop returns the last element and removes it from the list
// returns zero value if list is empty
func (l *MYTYPEList) Pop() MYTYPE {
	if len(l.elements) == 0 {
		return *new(MYTYPE)
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

// Grep iterates over the slice by calling the given function
// providing the slice index and the value
// and returning a new List that contains all element where
// the given function returned true
func (l *MYTYPEList) Grep(f func(int, MYTYPE) bool) MYTYPEList {
	result := MYTYPEList{}
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
func (l *MYTYPEList) Map(f func(int, MYTYPE) MYTYPE) MYTYPEList {
	result := MYTYPEList{}
	for index, value := range l.elements {
		result.elements = append(result.elements, f(index, value))
	}
	return result
}

// Any iterates over the slice by calling the given function
// providing the slice index and the value
// and returning true if one of the function results is true
// else false
func (l *MYTYPEList) Any(f func(int, MYTYPE) bool) bool {
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
func (l *MYTYPEList) All(f func(int, MYTYPE) bool) bool {
	for index, value := range l.elements {
		if f(index, value) == false {
			return false
		}
	}
	return true
}

// Reduce iterates over the slice by calling the given function
// providing the cumulated value, the slice index, and the value
func (l *MYTYPEList) Reduce(f func(MYTYPE, int, MYTYPE) MYTYPE) MYTYPE {
	var cum MYTYPE
	for index, value := range l.elements {
		if index == 0 {
			cum = value
		} else {
			cum = f(cum, index, value)
		}
	}
	return cum
}
