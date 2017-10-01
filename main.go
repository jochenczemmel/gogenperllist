package main

import (
	"fmt"
	"os"
	"strings"
	"text/template"
)

var code = `package {{.Package}}

type {{.MyType}}List struct {
	elements []{{.MyType}}
}

// New{{.MyType}}List creates a new list with optional initial values
func New{{.MyType}}List(values ...{{.MyType}}) *{{.MyType}}List {
	return &{{.MyType}}List{values}
}

// return len of list
func (l *{{.MyType}}List) Len() int {
	return len(l.elements)
}

// return cap of list
func (l *{{.MyType}}List) Cap() int {
	return cap(l.elements)
}

// append element to list
func (l *{{.MyType}}List) Push(element {{.MyType}}) {
	l.elements = append(l.elements, element)
}

// prepend element to list
func (l *{{.MyType}}List) Unshift(element {{.MyType}}) {
	newl := append([]{{.MyType}}{element}, l.elements...)
	l.elements = newl
}

// get first element of list
// or zero value if list is empty
func (l *{{.MyType}}List) Shift() {{.MyType}} {
	if len(l.elements) == 0 {
		r := new({{.MyType}})
		return *r
	}
	first := l.elements[0]
	l.elements = l.elements[1:]
	return first
}

// get last element of list
// or zero value if list is empty
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

// iterate over list elements
func (l *{{.MyType}}List) Foreach(f func(int, {{.MyType}})) {
	for index, value := range l.elements {
		f(index, value)
	}
}

// get last element of list
// true if list is not empty, else false
func (l *{{.MyType}}List) PopChecked() ({{.MyType}}, bool) {
	if len(l.elements) == 0 {
		r := new({{.MyType}})
		return *r, false
	}
	return l.Pop(), true
}

// get first element of list
// true if list is not empty, else false
func (l *{{.MyType}}List) ShiftChecked() ({{.MyType}}, bool) {
	if len(l.elements) == 0 {
		r := new({{.MyType}})
		return *r, false
	}
	return l.Shift(), true
}

// get element at index
// no error handling, default slice behaviour!
func (l *{{.MyType}}List) ElementAt(index int) {{.MyType}} {
	return l.elements[index]
}
`

func main() {
	parsedTemplate := template.Must(template.New("perllist").Parse(code))

	for i := 1; i < len(os.Args); i++ {
		filename := strings.TrimSuffix(os.Getenv("GOFILE"), ".go") + "_" +
			os.Args[i] + "_perllist_generated.go"
		fid, err := os.Create(filename)
		if err != nil {
			fmt.Fprintf(os.Stderr, "ERROR: could not create file %s: %s\n",
				filename, err)
			continue
		}
		parameter := map[string]string{
			"MyType":  os.Args[i],
			"Package": os.Getenv("GOPACKAGE"),
		}
		parsedTemplate.Execute(fid, parameter)
		fid.Close()
	}
}
