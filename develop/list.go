/*
Package develop implements some perl list functions for int lists
*/
package develop

type intList struct {
	elements []int
}

// NewIntList creates a new list with optional initial values
func NewIntList(values ...int) *intList {
	return &intList{values}
}

// return len of list
func (l *intList) Len() int {
	return len(l.elements)
}

// return cap of list
func (l *intList) Cap() int {
	return cap(l.elements)
}

// append element to list
func (l *intList) Push(element int) {
	l.elements = append(l.elements, element)
}

// prepend element to list
func (l *intList) Unshift(element int) {
	newl := append([]int{element}, l.elements...)
	l.elements = newl
}

// get first element of list
// or zero value if list is empty
func (l *intList) Shift() int {
	if len(l.elements) == 0 {
		r := new(int)
		return *r
	}
	first := l.elements[0]
	l.elements = l.elements[1:]
	return first
}

// get last element of list
// or zero value if list is empty
func (l *intList) Pop() int {
	if len(l.elements) == 0 {
		r := new(int)
		return *r
	}
	lastIndex := len(l.elements) - 1
	last := l.elements[lastIndex]
	l.elements = l.elements[:lastIndex]
	return last
}

// iterate over list elements
func (l *intList) Foreach(f func(int, int)) {
	for index, value := range l.elements {
		f(index, value)
	}
}

// get last element of list
// true if list is not empty, else false
func (l *intList) PopChecked() (int, bool) {
	if len(l.elements) == 0 {
		r := new(int)
		return *r, false
	}
	return l.Pop(), true
}

// get first element of list
// true if list is not empty, else false
func (l *intList) ShiftChecked() (int, bool) {
	if len(l.elements) == 0 {
		r := new(int)
		return *r, false
	}
	return l.Shift(), true
}

// get element at index
// no error handling, default slice behaviour!
func (l *intList) ElementAt(index int) int {
	return l.elements[index]
}
