package main

var code = `package {{.Package}}

type {{.MyType}}List struct {
	elements []{{.MyType}}
}

func New{{.MyType}}List(values ...{{.MyType}}) *{{.MyType}}List {
	return &{{.MyType}}List{values}
}

func (l *{{.MyType}}List) Len() int {
	return len(l.elements)
}

func (l *{{.MyType}}List) Cap() int {
	return cap(l.elements)
}

func (l *{{.MyType}}List) Push(element {{.MyType}}) {
	l.elements = append(l.elements, element)
}

func (l *{{.MyType}}List) Unshift(element {{.MyType}}) {
	newl := append([]{{.MyType}}{element}, l.elements...)
	l.elements = newl
}

func (l *{{.MyType}}List) Shift() {{.MyType}} {
	if len(l.elements) == 0 {
		r := new({{.MyType}})
		return *r
	}
	first := l.elements[0]
	l.elements = l.elements[1:]
	return first
}

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

func (l *{{.MyType}}List) Foreach(f func(int, {{.MyType}})) {
	for index, value := range l.elements {
		f(index, value)
	}
}

func (l *{{.MyType}}List) PopChecked() ({{.MyType}}, bool) {
	if len(l.elements) == 0 {
		r := new({{.MyType}})
		return *r, false
	}
	return l.Pop(), true
}

func (l *{{.MyType}}List) ShiftChecked() ({{.MyType}}, bool) {
	if len(l.elements) == 0 {
		r := new({{.MyType}})
		return *r, false
	}
	return l.Shift(), true
}

func (l *{{.MyType}}List) ElementAt(index int) {{.MyType}} {
	return l.elements[index]
}

`