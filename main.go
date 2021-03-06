package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

var parsedTemplate = template.Must(template.New("perllist").Parse(code))

func main() {
	// parse the template generated by develop/develop

	importDir := flag.String("import", "", "import path for type")
	flag.Parse()
	prefix := extractPrefix(*importDir)

	// genereate code for all given types
	for _, arg := range flag.Args() {
		printFile(arg, prefix, *importDir)
	}

}

func printFile(typename, prefix, importDir string) {

	// get file name where go:generate is located
	// create generated file name
	filename := strings.TrimSuffix(os.Getenv("GOFILE"), ".go") + "_" +
		typename + "_perllist_generated.go"

	fid, err := os.Create(filename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "ERROR: could not create file %s: %s\n",
			filename, err)
		os.Exit(2)
	}
	// create values to fill in the template
	parameter := map[string]string{
		"Package":     os.Getenv("GOPACKAGE"),
		"Import":      importDir,
		"MyTypeShort": typename,
		"MyType":      prefix + typename,
	}
	// replace the template values
	parsedTemplate.Execute(fid, parameter)
	fid.Close()
}

func extractPrefix(in string) string {
	if in == "" {
		return ""
	}
	beforeDot := strings.Split(filepath.Base(in), ".")
	if strings.Index(beforeDot[0], "-") > -1 {
		afterDash := strings.Split(beforeDot[0], "-")
		return afterDash[1] + "."
	}
	return beforeDot[0] + "."
}
