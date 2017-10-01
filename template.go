package main

var code = `package {{.Package}}

// {{.MyType}}List contains a slice of {{.MyType}} values
type {{.MyType}}List struct {
	elements []{{.MyType}}
}

// New{{.MyType}}List creates a new List of {{.MyType}} values
// with optional initial values
func New{{.MyType}}List(values ...{{.MyType}}) *{{.MyType}}List {
	return &{{.MyType}}List{values}
}

// Len returns the length of the underlying slice
func (l *{{.MyType}}List) Len() int {
	return len(l.elements)
}

// Cap returns the capacity of the underlying slice
func (l *{{.MyType}}List) Cap() int {
	return cap(l.elements)
}

// Push appends an element at the end of the list
func (l *{{.MyType}}List) Push(element {{.MyType}}) {
	l.elements = append(l.elements, element)
}

// Unshift adds an element at the beginning of the list
func (l *{{.MyType}}List) Unshift(element {{.MyType}}) {
	newl := append([]{{.MyType}}{element}, l.elements...)
	l.elements = newl
}

// Shift returns the first element and removes it from the list
// returns zero value if list is empty
func (l *{{.MyType}}List) Shift() {{.MyType}} {
	if len(l.elements) == 0 {
		r := new({{.MyType}})
		return *r
	}
	first := l.elements[0]
	l.elements = l.elements[1:]
	return first
}

// Pop returns the last element and removes it from the list
// returns zero value if list is empty
func (l *{{.MyType}}List) Pop() {{.MyType}} {
	if len(l.elements) == 0 {
		r := new({{.MyType}})
		return *r
	}
	lastIndex := len(l.elements) - 1
	last := l.elements[lastIndex]
	l.elements = l.elements[:lastIndex]
	return last
}

// Foreach iterates over the slice by calling the given function
// providing the slice index and the value
func (l *{{.MyType}}List) Foreach(f func(int, {{.MyType}})) {
	for index, value := range l.elements {
		f(index, value)
	}
}

// Pop returns the last element and removes it from the list
// returns the value and true if the list contains values,
// else the zero value and false
func (l *{{.MyType}}List) PopChecked() ({{.MyType}}, bool) {
	if len(l.elements) == 0 {
		r := new({{.MyType}})
		return *r, false
	}
	return l.Pop(), true
}

// Shift returns the first element and removes it from the list
// returns the value and true if the list contains values,
// else the zero value and false
func (l *{{.MyType}}List) ShiftChecked() ({{.MyType}}, bool) {
	if len(l.elements) == 0 {
		r := new({{.MyType}})
		return *r, false
	}
	return l.Shift(), true
}

// ElementAt returns the element at the given position.
// It panics if index is invalid (default slice behaviour)
func (l *{{.MyType}}List) ElementAt(index int) {{.MyType}} {
	return l.elements[index]
}

`