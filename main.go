package main

import (
	"fmt"
	"os"
	"strings"
	"text/template"
)

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
