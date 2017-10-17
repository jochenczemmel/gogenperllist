package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
)

func main() {

	content := readFile("list.go")
	newcontent := replace(content)
	write(newcontent, "../template.go")
}

func write(code, filename string) {

	fid, err := os.Create("../template.go")
	if err != nil {
		fmt.Fprintf(os.Stderr, "ERROR: could not create file %s: %s\n",
			filename, err)
		os.Exit(2)
	}

	defer fid.Close()

	code = "package main\n\nvar code = `" + code + "\n`"

	_, err = fid.Write([]byte(code))
	if err != nil {
		fmt.Fprintf(os.Stderr, "ERROR: could not write file %s: %s\n",
			filename, err)
		os.Exit(2)
	}

}

func replace(in string) string {
	out := regexp.MustCompile("package main").ReplaceAllString(in,
		"package {{.Package}}")
	out = regexp.MustCompile("// REPlACEIMPORT").ReplaceAllString(out,
		`{{if .Import}}import "{{.Import}}" {{end}}`)
	out = regexp.MustCompile("NewMYTYPE").ReplaceAllString(out,
		"New{{.MyTypeShort}}")
	out = regexp.MustCompile("MYTYPEList").ReplaceAllString(out,
		"{{.MyTypeShort}}List")

	return regexp.MustCompile("MYTYPE").ReplaceAllString(out,
		"{{.MyType}}")
}

func readFile(filename string) string {

	content, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "ERROR: could not read file %s: %s\n",
			filename, err)
		os.Exit(2)
	}

	return string(content)
}
