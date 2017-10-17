package main



// PlayerList contains a slice of Player values
type PlayerList struct {
	elements []Player
}

// NewPlayerList creates a new List of Player values
// with optional initial values
func NewPlayerList(values ...Player) *PlayerList {
	return &PlayerList{values}
}

// Len returns the length of the underlying slice
func (l *PlayerList) Len() int {
	return len(l.elements)
}

// Cap returns the capacity of the underlying slice
func (l *PlayerList) Cap() int {
	return cap(l.elements)
}

// Push appends an element at the end of the list
func (l *PlayerList) Push(element Player) {
	l.elements = append(l.elements, element)
}

// Unshift adds an element at the beginning of the list
func (l *PlayerList) Unshift(element Player) {
	newl := append([]Player{element}, l.elements...)
	l.elements = newl
}

// Shift returns the first element and removes it from the list
// returns zero value if list is empty
func (l *PlayerList) Shift() Player {
	if len(l.elements) == 0 {
		return *new(Player)
	}
	first := l.elements[0]
	l.elements = l.elements[1:]
	return first
}

// Pop returns the last element and removes it from the list
// returns zero value if list is empty
func (l *PlayerList) Pop() Player {
	if len(l.elements) == 0 {
		return *new(Player)
	}
	lastIndex := len(l.elements) - 1
	last := l.elements[lastIndex]
	l.elements = l.elements[:lastIndex]
	return last
}

// Foreach iterates over the slice by calling the given function
// providing the slice index and the value
func (l *PlayerList) Foreach(f func(int, Player)) {
	for index, value := range l.elements {
		f(index, value)
	}
}

// PopChecked returns the last element and removes it from the list
// returns the value and true if the list contains values,
// else the zero value and false
func (l *PlayerList) PopChecked() (Player, bool) {
	if len(l.elements) == 0 {
		r := new(Player)
		return *r, false
	}
	return l.Pop(), true
}

// ShiftChecked returns the first element and removes it from the list
// returns the value and true if the list contains values,
// else the zero value and false
func (l *PlayerList) ShiftChecked() (Player, bool) {
	if len(l.elements) == 0 {
		r := new(Player)
		return *r, false
	}
	return l.Shift(), true
}

// ElementAt returns the element at the given position.
// It panics if index is invalid (default slice behaviour)
func (l *PlayerList) ElementAt(index int) Player {
	return l.elements[index]
}

// Grep iterates over the slice by calling the given function
// providing the slice index and the value
// and returning a new List that contains all element where
// the given function returned true
func (l *PlayerList) Grep(f func(int, Player) bool) PlayerList {
	result := PlayerList{}
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
func (l *PlayerList) Map(f func(int, Player) Player) PlayerList {
	result := PlayerList{}
	for index, value := range l.elements {
		result.elements = append(result.elements, f(index, value))
	}
	return result
}

// Any iterates over the slice by calling the given function
// providing the slice index and the value
// and returning true if one of the function results is true
// else false
func (l *PlayerList) Any(f func(int, Player) bool) bool {
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
func (l *PlayerList) All(f func(int, Player) bool) bool {
	for index, value := range l.elements {
		if f(index, value) == false {
			return false
		}
	}
	return true
}

// Fold iterates over the slice by calling the given function
// providing the cumulated value, the slice index, and the value
func (l *PlayerList) Fold(f func(Player, int, Player) Player) Player {
	var cum Player
	for index, value := range l.elements {
		cum = f(cum, index, value)
	}
	return cum
}

