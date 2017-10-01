package main

type PlayerList struct {
	elements []Player
}

func NewPlayerList(values ...Player) *PlayerList {
	return &PlayerList{values}
}

func (l *PlayerList) Len() int {
	return len(l.elements)
}

func (l *PlayerList) Cap() int {
	return cap(l.elements)
}

func (l *PlayerList) Push(element Player) {
	l.elements = append(l.elements, element)
}

func (l *PlayerList) Unshift(element Player) {
	newl := append([]Player{element}, l.elements...)
	l.elements = newl
}

func (l *PlayerList) Shift() Player {
	if len(l.elements) == 0 {
		r := new(Player)
		return *r
	}
	first := l.elements[0]
	l.elements = l.elements[1:]
	return first
}

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

func (l *PlayerList) Foreach(f func(int, Player)) {
	for index, value := range l.elements {
		f(index, value)
	}
}

func (l *PlayerList) PopChecked() (Player, bool) {
	if len(l.elements) == 0 {
		r := new(Player)
		return *r, false
	}
	return l.Pop(), true
}

func (l *PlayerList) ShiftChecked() (Player, bool) {
	if len(l.elements) == 0 {
		r := new(Player)
		return *r, false
	}
	return l.Shift(), true
}

func (l *PlayerList) ElementAt(index int) Player {
	return l.elements[index]
}

