package main

import (
	"fmt"
	"os"
	"strings"
	"text/template"
)

func main() {
	// parse the template generated by develop/develop
	parsedTemplate := template.Must(template.New("perllist").Parse(code))

	// genereate code for all given types
	for i := 1; i < len(os.Args); i++ {
		// get file name where go:generate is located
		// create generated file name
		filename := strings.TrimSuffix(os.Getenv("GOFILE"), ".go") + "_" +
			os.Args[i] + "_perllist_generated.go"

		fid, err := os.Create(filename)
		if err != nil {
			fmt.Fprintf(os.Stderr, "ERROR: could not create file %s: %s\n",
				filename, err)
			os.Exit(2)
		}
		// create values to fill in the template
		parameter := map[string]string{
			"MyType":  os.Args[i],
			"Package": os.Getenv("GOPACKAGE"),
		}
		// replace the template values
		parsedTemplate.Execute(fid, parameter)
		fid.Close()
	}
}
