package main

type MYTYPEList struct {
	elements []MYTYPE
}

func NewMYTYPEList(values ...MYTYPE) *MYTYPEList {
	return &MYTYPEList{values}
}

func (l *MYTYPEList) Len() int {
	return len(l.elements)
}

func (l *MYTYPEList) Cap() int {
	return cap(l.elements)
}

func (l *MYTYPEList) Push(element MYTYPE) {
	l.elements = append(l.elements, element)
}

func (l *MYTYPEList) Unshift(element MYTYPE) {
	newl := append([]MYTYPE{element}, l.elements...)
	l.elements = newl
}

func (l *MYTYPEList) Shift() MYTYPE {
	if len(l.elements) == 0 {
		r := new(MYTYPE)
		return *r
	}
	first := l.elements[0]
	l.elements = l.elements[1:]
	return first
}

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

func (l *MYTYPEList) Foreach(f func(int, MYTYPE)) {
	for index, value := range l.elements {
		f(index, value)
	}
}

func (l *MYTYPEList) PopChecked() (MYTYPE, bool) {
	if len(l.elements) == 0 {
		r := new(MYTYPE)
		return *r, false
	}
	return l.Pop(), true
}

func (l *MYTYPEList) ShiftChecked() (MYTYPE, bool) {
	if len(l.elements) == 0 {
		r := new(MYTYPE)
		return *r, false
	}
	return l.Shift(), true
}

func (l *MYTYPEList) ElementAt(index int) MYTYPE {
	return l.elements[index]
}
