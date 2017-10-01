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

func write(code []byte, filename string) {

	fid, err := os.Create("../template.go")
	if err != nil {
		fmt.Fprintf(os.Stderr, "ERROR: could not create file %s: %s\n",
			filename, err)
		os.Exit(2)
	}

	defer fid.Close()

	_, err = fid.Write(code)
	if err != nil {
		fmt.Fprintf(os.Stderr, "ERROR: could not write file %s: %s\n",
			filename, err)
		os.Exit(2)
	}

}

func replace(in []byte) []byte {
	var out []byte
	regexp.MustCompile("MYTYPE").ReplaceAll(in, out)
	return out
}

func readFile(filename string) []byte {

	content, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "ERROR: could not read file %s: %s\n",
			filename, err)
		os.Exit(2)
	}

	return content
}
