package main

var code = `package {{.Package}}

{{if .Import}}import "{{.Import}}" {{end}}

// {{.MyTypeShort}}List contains a slice of {{.MyType}} values
type {{.MyTypeShort}}List struct {
	elements []{{.MyType}}
}

// New{{.MyTypeShort}}List creates a new List of {{.MyType}} values
// with optional initial values
func New{{.MyTypeShort}}List(values ...{{.MyType}}) *{{.MyTypeShort}}List {
	return &{{.MyTypeShort}}List{values}
}

// Len returns the length of the underlying slice
func (l *{{.MyTypeShort}}List) Len() int {
	return len(l.elements)
}

// Cap returns the capacity of the underlying slice
func (l *{{.MyTypeShort}}List) Cap() int {
	return cap(l.elements)
}

// Push appends an element at the end of the list
func (l *{{.MyTypeShort}}List) Push(element {{.MyType}}) {
	l.elements = append(l.elements, element)
}

// Unshift adds an element at the beginning of the list
func (l *{{.MyTypeShort}}List) Unshift(element {{.MyType}}) {
	newl := append([]{{.MyType}}{element}, l.elements...)
	l.elements = newl
}

// Shift returns the first element and removes it from the list
// returns zero value if list is empty
func (l *{{.MyTypeShort}}List) Shift() {{.MyType}} {
	if len(l.elements) == 0 {
		return *new({{.MyType}})
	}
	first := l.elements[0]
	l.elements = l.elements[1:]
	return first
}

// Pop returns the last element and removes it from the list
// returns zero value if list is empty
func (l *{{.MyTypeShort}}List) Pop() {{.MyType}} {
	if len(l.elements) == 0 {
		return *new({{.MyType}})
	}
	lastIndex := len(l.elements) - 1
	last := l.elements[lastIndex]
	l.elements = l.elements[:lastIndex]
	return last
}

// Foreach iterates over the slice by calling the given function
// providing the slice index and the value
func (l *{{.MyTypeShort}}List) Foreach(f func(int, {{.MyType}})) {
	for index, value := range l.elements {
		f(index, value)
	}
}

// PopChecked returns the last element and removes it from the list
// returns the value and true if the list contains values,
// else the zero value and false
func (l *{{.MyTypeShort}}List) PopChecked() ({{.MyType}}, bool) {
	if len(l.elements) == 0 {
		r := new({{.MyType}})
		return *r, false
	}
	return l.Pop(), true
}

// ShiftChecked returns the first element and removes it from the list
// returns the value and true if the list contains values,
// else the zero value and false
func (l *{{.MyTypeShort}}List) ShiftChecked() ({{.MyType}}, bool) {
	if len(l.elements) == 0 {
		r := new({{.MyType}})
		return *r, false
	}
	return l.Shift(), true
}

// ElementAt returns the element at the given position.
// It panics if index is invalid (default slice behaviour)
func (l *{{.MyTypeShort}}List) ElementAt(index int) {{.MyType}} {
	return l.elements[index]
}

// Grep iterates over the slice by calling the given function
// providing the slice index and the value
// and returning a new List that contains all element where
// the given function returned true
func (l *{{.MyTypeShort}}List) Grep(f func(int, {{.MyType}}) bool) {{.MyTypeShort}}List {
	result := {{.MyTypeShort}}List{}
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
func (l *{{.MyTypeShort}}List) Map(f func(int, {{.MyType}}) {{.MyType}}) {{.MyTypeShort}}List {
	result := {{.MyTypeShort}}List{}
	for index, value := range l.elements {
		result.elements = append(result.elements, f(index, value))
	}
	return result
}

// Any iterates over the slice by calling the given function
// providing the slice index and the value
// and returning true if one of the function results is true
// else false
func (l *{{.MyTypeShort}}List) Any(f func(int, {{.MyType}}) bool) bool {
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
func (l *{{.MyTypeShort}}List) All(f func(int, {{.MyType}}) bool) bool {
	for index, value := range l.elements {
		if f(index, value) == false {
			return false
		}
	}
	return true
}

// Fold iterates over the slice by calling the given function
// providing the cumulated value, the slice index, and the value
func (l *{{.MyTypeShort}}List) Fold(f func({{.MyType}}, int, {{.MyType}}) {{.MyType}}) {{.MyType}} {
	var cum {{.MyType}}
	for index, value := range l.elements {
		cum = f(cum, index, value)
	}
	return cum
}

`