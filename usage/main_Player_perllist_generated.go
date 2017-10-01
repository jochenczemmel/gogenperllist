package main

type PlayerList struct {
	elements []Player
}

// NewPlayerList creates a new list with optional initial values
func NewPlayerList(values ...Player) *PlayerList {
	return &PlayerList{values}
}

// return len of list
func (l *PlayerList) Len() int {
	return len(l.elements)
}

// return cap of list
func (l *PlayerList) Cap() int {
	return cap(l.elements)
}

// append element to list
func (l *PlayerList) Push(element Player) {
	l.elements = append(l.elements, element)
}

// prepend element to list
func (l *PlayerList) Unshift(element Player) {
	newl := append([]Player{element}, l.elements...)
	l.elements = newl
}

// get first element of list
// or zero value if list is empty
func (l *PlayerList) Shift() Player {
	if len(l.elements) == 0 {
		r := new(Player)
		return *r
	}
	first := l.elements[0]
	l.elements = l.elements[1:]
	return first
}

// get last element of list
// or zero value if list is empty
func (l *PlayerList) Pop() Player {
	if len(l.elements) == 0 {
		r := new(Player)
		return *r
	}
	lastIndex := len(l.elements) - 1
	last := l.elements[lastIndex]
	l.elements = l.elements[:lastIndex]
	return last
}

// iterate over list elements
func (l *PlayerList) Foreach(f func(int, Player)) {
	for index, value := range l.elements {
		f(index, value)
	}
}

// get last element of list
// true if list is not empty, else false
func (l *PlayerList) PopChecked() (Player, bool) {
	if len(l.elements) == 0 {
		r := new(Player)
		return *r, false
	}
	return l.Pop(), true
}

// get first element of list
// true if list is not empty, else false
func (l *PlayerList) ShiftChecked() (Player, bool) {
	if len(l.elements) == 0 {
		r := new(Player)
		return *r, false
	}
	return l.Shift(), true
}

// get element at index
// no error handling, default slice behaviour!
func (l *PlayerList) ElementAt(index int) Player {
	return l.elements[index]
}
