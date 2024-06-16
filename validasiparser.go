package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

type HTMLParser struct {
	stack []string
}

func NewHTMLParser() *HTMLParser {
	return &HTMLParser{stack: []string{}}
}

func (parser *HTMLParser) push(tag string) {
	parser.stack = append(parser.stack, tag)
}

func (parser *HTMLParser) pop() (string, bool) {
	if len(parser.stack) == 0 {
		return "", false
	}
	tag := parser.stack[len(parser.stack)-1]
	parser.stack = parser.stack[:len(parser.stack)-1]
	return tag, true
}

func (parser *HTMLParser) isEmpty() bool {
	return len(parser.stack) == 0
}

func (parser *HTMLParser) isValidHTML(html string) bool {
	parser.stack = []string{} // Kosongkan stack di awal
	i := 0
	for i < len(html) {
		if html[i] == '<' {
			// mencari tag penutup
			j := i + 1
			for j < len(html) && html[j] != '>' {
				j++
			}
			if j == len(html) {
				return false // tag gk valid
			}

			tag := html[i+1 : j]
			if tag[0] != '/' {
				// memasukan opening ke stack
				parser.push(tag)
			} else {
				openingTag, ok := parser.pop()
				if !ok || openingTag != tag[1:] {
					return false
				}
			}
			i = j
		}
		i++
	}
	// validasi stack kosong
	return parser.isEmpty()
}

func (parser *HTMLParser) isValidHTMLStructure(html string) bool {
	if !parser.isValidHTML(html) {
		return false
	}
	// memastikan file html memiliki root
	if !strings.Contains(html, "<html>") || !strings.Contains(html, "</html>") {
		return false
	}
	// memastikan tag <head> dan <body> benar di dalam <html>
	headStart := strings.Index(html, "<head>")
	headEnd := strings.Index(html, "</head>")
	bodyStart := strings.Index(html, "<body>")
	bodyEnd := strings.Index(html, "</body>")

	if headStart != -1 && (headEnd == -1 || headEnd < headStart) {
		return false
	}
	if bodyStart != -1 && (bodyEnd == -1 || bodyEnd < bodyStart) {
		return false
	}
	if headEnd != -1 && bodyStart != -1 && headEnd > bodyStart {
		return false
	}

	return true
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run validasiparser.go <html_file1> [<html_file2> ...]")
		return
	}

	parser := NewHTMLParser()

	for _, filename := range os.Args[1:] {
		htmlBytes, err := ioutil.ReadFile(filename)
		if err != nil {
			fmt.Printf("Error reading file %s: %s\n", filename, err)
			continue
		}

		html := string(htmlBytes)
		result := parser.isValidHTMLStructure(html)
		status := "Rejected"
		if result {
			status = "Accepted"
		}
		fmt.Printf("File: %s | Status: %s\n", filename, status)
	}
}
