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

func New{{.MyType}}List() *{{.MyType}}List {
	return &{{.MyType}}List{[]{{.MyType}}{}}
}

func (liste *{{.MyType}}List) Len() int {
	return len(liste.elements)
}

func (liste *{{.MyType}}List) Push(element {{.MyType}}) {
	liste.elements = append(liste.elements, element)
}

func (liste *{{.MyType}}List) Unshift(element {{.MyType}}) {
	newl := make([]{{.MyType}}, 0, len(liste.elements)+1)
	newl = append(newl, element)
	newl = append(newl, liste.elements...)
	liste.elements = newl
}

func (liste *{{.MyType}}List) Shift() {{.MyType}} {
	if len(liste.elements) == 0 {
		r := new({{.MyType}})
		return *r
	}
	first := liste.elements[0]
	liste.elements = liste.elements[1:]
	return first
}

func (liste *{{.MyType}}List) Pop() {{.MyType}} {
	if len(liste.elements) == 0 {
		r := new({{.MyType}})
		return *r
	}
	lastIndex := len(liste.elements) - 1
	last := liste.elements[lastIndex]
	liste.elements = liste.elements[:lastIndex]
	return last
}

func (liste *{{.MyType}}List) Foreach(f func(int, {{.MyType}})) {
	for index, value := range liste.elements {
		f(index, value)
	}
}`

func main() {
	parsedTemplate := template.Must(template.New("liste").Parse(code))

	for i := 1; i < len(os.Args); i++ {
		filename := strings.TrimSuffix(os.Getenv("GOFILE"), ".go") + "_" +
			strings.ToLower(os.Args[i]) + "_liste_generated.go"
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
